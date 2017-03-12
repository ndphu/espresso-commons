package messaging

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
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
	OnNewMessage(*Message)
}

type ListenerRecord struct {
	topic    string
	listener MessageListener
}

func (m *MessageRouter) Stop() {
	m.running = false
}

func (m *MessageRouter) GetMQTTClient() mqtt.Client {
	return m.client
}

func (m *MessageRouter) loop() {
	for m.running {
		if !m.client.IsConnected() {
			log.Println("Trying to reconnect to the broker...")
			if token := m.client.Connect(); token.Wait() && token.Error() != nil {
				log.Println("Faild to reconnect")
			} else {
				log.Println("Reconnected to the broker")
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
	msg := Message{}
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

func (m *MessageRouter) Publish(msg Message) error {
	if m.client.IsConnected() {
		data, err := json.Marshal(msg)
		if err != nil {
			return err
		}

		rawMsg := string(data)
		log.Println("Publish", rawMsg)
		if token := m.client.Publish(string(msg.Destination), DefaultPubQos, false, rawMsg); token.Wait() && token.Error() != nil {
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
	} else {
		log.Println("Subscribed to topic", topic)
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
