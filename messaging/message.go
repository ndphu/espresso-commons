package messaging

type MessageType string
type MessageSource string
type MessageDestination string

var (
	MessageType_IREventAdded       MessageType        = "IR_EVENT_ADDED"
	MessageType_ExecuteTextCommand MessageType        = "EXEC_TEXT_COMMAND"
	MessageType_ExecuteGPIOCommand MessageType        = "EXEC_GPIO_COMMAND"
	MessageSource_UI               MessageSource      = "ui"
	MessageSource_IR_AGENT         MessageSource      = "ir-agent"
	MessageDestination_TextCommand MessageDestination = "/espresso/components/commands/text_commands"
	MessageDestination_GPIOCommand MessageDestination = "/espresso/components/commands/gpio_commands"
)

type Message struct {
	Destination MessageDestination `json:"destination"`
	Source      MessageSource      `json:"source"`
	Type        MessageType        `json:"type"`
	Payload     string             `json:"payload"`
}
