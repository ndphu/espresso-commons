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
	CommandTopicTemplate        = "/espresso/devices/%s/commands"
	DefaultToDeviceQos   byte   = 1
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
