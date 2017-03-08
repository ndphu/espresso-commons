package messaging

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/ndphu/espresso-commons/model"
	"log"
	"time"
)

var (
	DefaultPubQos byte = 0
	DefaultSubQos byte = 0
)

type MessageRouter struct {
	client    mqtt.Client
	listeners map[string]([]MessageListener)
	running   bool
	//listenerRecords []ListenerRecord
}

type MessageListener interface {
	OnNewMessage(*model.Message)
}

type ListenerRecord struct {
	topic    string
	listener MessageListener
}

func (m *MessageRouter) Stop() {
	m.running = false
}

func (m *MessageRouter) loop() {
	for m.running {
		if !m.client.IsConnected() {
			log.Println("Trying to reconnect to the broker...")
			if token := m.client.Connect(); token.Wait() && token.Error() != nil {
				log.Println("Faild to reconnect")
			} else {
				log.Println("Reconnected to the broker")
				// subscribe current subscriber
				// lc := len(m.listenerRecords)
				// tmp:=make([])
				// m.listenerRecords = make([]ListenerRecord, 0)
				// for i := 0; i < len(tmp); i++ {
				// 	m.Subscribe(tmp[i].topic, tmp[i].listener)
				// }
				// for i := 0; i < len(m.listenerRecords); i++ {
				// 	m.Subscribe(m.listenerRecords[i].topic, m.listenerRecords[i].listener)
				// }
				for k, _ := range m.listeners {
					token := m.client.Subscribe(k, DefaultSubQos, func(c mqtt.Client, msg mqtt.Message) {
						m.HandleMessage(c, msg)
					})
					token.Wait()
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func (m *MessageRouter) HandleMessage(client mqtt.Client, mqttMessage mqtt.Message) {
	topic := mqttMessage.Topic()
	ls := m.listeners[topic]
	msg := model.Message{}
	json.Unmarshal(mqttMessage.Payload(), &msg)
	for i := 0; i < len(ls); i++ {
		go ls[i].OnNewMessage(&msg)
	}
}

func NewMessageRouter(h string, p int, u string, pwd string, clientId string) (*MessageRouter, error) {
	opt := mqtt.NewClientOptions()
	opt.AddBroker(fmt.Sprintf("tcp://%s:%d", h, p))
	opt.SetUsername(u)
	opt.SetPassword(pwd)
	// manual manage reconnect
	opt.SetAutoReconnect(false)
	if len(clientId) == 0 {
		return nil, errors.New("Empty client id")
	}
	opt.SetClientID(clientId)

	c := mqtt.NewClient(opt)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	router := &MessageRouter{
		client:    c,
		running:   true,
		listeners: make(map[string][]MessageListener),
	}

	go router.loop()

	return router, nil

}

func (m *MessageRouter) Publish(msg model.Message) error {
	if m.client.IsConnected() {
		data, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		if token := m.client.Publish(msg.Destination, DefaultPubQos, false, string(data)); token.Wait() && token.Error() != nil {
			return token.Error()
		}
		return nil
	} else {
		return errors.New("MQTT client disconnected")
	}
}

func (m *MessageRouter) Subscribe(topic string, listener MessageListener) error {
	if len(topic) == 0 {
		return errors.New("Empty topic field")
	}
	if token := m.client.Subscribe(topic, DefaultSubQos, func(c mqtt.Client, msg mqtt.Message) {
		m.HandleMessage(c, msg)
	}); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	m.listeners[topic] = append(m.listeners[topic], listener)
	// m.listenerRecords = append(m.listenerRecords, ListenerRecord{
	// 	topic:    topic,
	// 	listener: listener,
	// })
	return nil
}

func (m *MessageRouter) Unsubscribe(topic string, listener MessageListener) error {
	if len(topic) == 0 {
		return errors.New("Empty topic field")
	}
	if token := m.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
