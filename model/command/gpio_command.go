package command

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type GPIOCommand struct {
	Id             bson.ObjectId `json:"_id" bson:"_id"`
	Pin            int           `json:"pin" bson:"pin"`
	State          bool          `json:"state" bson:"state"`
	Timestamp      time.Time     `json:"timestamp" bson:"timestamp"`
	TargetDeviceId bson.ObjectId `json:"targetDeviceId" bson:"targetDeviceId"`
}

func (t *GPIOCommand) GetObjectId() bson.ObjectId {
	return t.Id
}

func (t *GPIOCommand) SetObjectId(id bson.ObjectId) {
	t.Id = id
}

// func TextCommandFromPayload(payload interface{}) TextCommand{
//     data := payload.(map[string]string)
//     return TextCommand{
//         Id:
//     }
// }
