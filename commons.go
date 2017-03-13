package commons

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	DBName               string = "espresso"
	IRAgentEventTopic           = "/espresso/components/ir-agent/events"
	TextCommandTopic            = "/espresso/components/device-manager/text-command"
	CommandTopicTemplate        = "/espresso/device/%s/commands"
	DefaultToDeviceQos   byte   = 1
	// MQTT
	BrokerHost string = "19november.freeddns.org"
	BrokerPort int    = 5370
)

func GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()

}

func GetRandomWithSeed(seed int64) int {
	rand.Seed(seed)
	return rand.Int()
}

func GetCommandTopicFromSerial(serial string) string {
	return fmt.Sprintf(CommandTopicTemplate, serial)
}
