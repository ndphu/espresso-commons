package command

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TextCommand struct {
	Id             bson.ObjectId `json:"_id" bson:"_id"`
	Text           string        `json:"text" bson:"text"`
	Timestamp      time.Time     `json:"timestamp" bson:"timestamp"`
	TargetDeviceId bson.ObjectId `json:"targetDeviceId" bson:"targetDeviceId"`
}

func (t *TextCommand) GetObjectId() bson.ObjectId {
	return t.Id
}

func (t *TextCommand) SetObjectId(id bson.ObjectId) {
	t.Id = id
}

// func TextCommandFromPayload(payload interface{}) TextCommand{
//     data := payload.(map[string]string)
//     return TextCommand{
//         Id:
//     }
// }
