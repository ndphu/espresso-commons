package model

type Message struct {
	Destination string      `json:"destination"`
	Source      string      `json:"source"`
	Type        string      `json:"type"`
	Payload     interface{} `json:"payload"`
}
