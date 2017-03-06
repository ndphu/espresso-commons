package messaging

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/ndphu/espresso-commons/model"
)

var (
	DefaultPubQos byte = 0
	DefaultSubQos byte = 0
)

type MessageRouter struct {
	client    mqtt.Client
	listeners map[string]([]MessageListener)
}

type MessageListener interface {
	OnNewMessage(*model.Message)
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
	if len(clientId) == 0 {
		return nil, errors.New("Empty client id")
	}
	opt.SetClientID(clientId)

	c := mqtt.NewClient(opt)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &MessageRouter{
		client:    c,
		listeners: make(map[string][]MessageListener),
	}, nil

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
