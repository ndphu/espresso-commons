package commons

import (
	"math/rand"
	"time"
)

var (
	DBName            string = "espresso"
	IRAgentEventTopic        = "components/ir-agent/events"
)

func GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()

}

func GetRandomWithSeed(seed int64) int {
	rand.Seed(seed)
	return rand.Int()
}
