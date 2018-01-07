package commons

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	DBName               string = "espresso"
	CommandTopicTemplate        = "/esp/device/%s/commands"
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

func GetEnv(name string, defaultValue string) string {
	value := os.Getenv(name)
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}

func RPiGetSerial(cpuInfoPath string) (s string, err error) {
	cpuInfo, e := os.Open(cpuInfoPath)
	if e != nil {
		return "", e
	}
	defer cpuInfo.Close()
	scanner := bufio.NewScanner(cpuInfo)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "Serial") == 0 {
			return strings.TrimSpace(strings.Split(line, ":")[1]), nil
		}
	}
	return "", errors.New("Fail to find serial")
}
