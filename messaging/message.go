package messaging

type MessageType string
type MessageSource string
type MessageDestination string

var (
	MessageType_IREventAdded         MessageType        = "IR_EVENT_ADDED"
	MessageType_ExecuteTextCommand   MessageType        = "EXEC_TEXT_COMMAND"
	MessageType_ExecuteGPIOCommand   MessageType        = "EXEC_GPIO_COMMAND"
	MessageType_DeviceOnline         MessageType        = "DEVICE_ONLINE"
	MessageType_DeviceOffline        MessageType        = "DEVICE_OFFLINE"
	MessageType_DeviceAdded          MessageType        = "DEVICE_ADDED"
	MessageType_DeviceRemoved        MessageType        = "DEVICE_REMOVED"
	MessageType_DeviceUpdated        MessageType        = "DEVICE_UPDATED"
	MessageSource_UI                 MessageSource      = "ui"
	MessageSource_IR_AGENT           MessageSource      = "ir-agent"
	MessageSource_DeviceManager      MessageSource      = "device-manager"
	MessageDestination_TextCommand   MessageDestination = "/espresso/components/commands/text_commands"
	MessageDestination_GPIOCommand   MessageDestination = "/espresso/components/commands/gpio_commands"
	MessageDestination_DeviceUpdated MessageDestination = "/espresso/components/devices/device_updated"
)

type Message struct {
	Destination MessageDestination `json:"destination"`
	Source      MessageSource      `json:"source"`
	Type        MessageType        `json:"type"`
	Payload     string             `json:"payload"`
}
