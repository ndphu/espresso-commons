package messaging

// Type of the message
type MessageType string

// Where the message come frome
type Source string

// The topic to publish message
type Topic string

var (
	// Message Type
	IREventAdded     MessageType = "IR_EVENT_ADDED"
	TextCommandAdded MessageType = "TEXT_COMMAND_ADDED"
	GPIOCommandAdded MessageType = "GPIO_COMMAND_ADDED"
	DeviceOnline     MessageType = "DEVICE_ONLINE"
	DeviceOffline    MessageType = "DEVICE_OFFLINE"
	DeviceAdded      MessageType = "DEVICE_ADDED"
	DeviceRemoved    MessageType = "DEVICE_REMOVED"
	DeviceUpdated    MessageType = "DEVICE_UPDATED"

	// Sourcce
	UI            Source = "ui"
	IRAgent       Source = "ir-agent"
	DeviceManager Source = "device-manager"

	// Topic
	IPCDevice  Topic = "/esp/ipc/device"
	IPCRule    Topic = "/esp/ipc/rule"
	IPCEvent   Topic = "/esp/ipc/event"
	IPCSensor  Topic = "/esp/ipc/sensor"
	IPCCommand Topic = "/esp/ipc/command"
)

// The data used to communicate between components
type Message struct {
	// The topic to publish this message
	Destination Topic `json:"destination"`
	// Indicate where the message come from
	// This is useful for consumer to decide if it need to process this message
	Source Source `json:"source"`
	// Kind of event for the consumer to process
	Type MessageType `json:"type"`
	// Message payload. Typically it is an object id related to the message type
	Payload string `json:"payload"`
}
