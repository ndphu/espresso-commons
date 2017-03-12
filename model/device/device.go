package device

import (
	"gopkg.in/mgo.v2/bson"
)

var (
	TypeDeviceAdd = "DEVICE_ADDED"
)

type Device struct {
	Id      bson.ObjectId `json:"_id" bson:"_id"`
	Name    string        `json:"name"`
	Serial  string        `json:"serial"`
	Managed bool          `json:"managed"`
}

func (ir *Device) GetObjectId() bson.ObjectId {
	return ir.Id
}

func (ir *Device) SetObjectId(id bson.ObjectId) {
	ir.Id = id
}
