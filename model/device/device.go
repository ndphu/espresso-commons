package device

import (
	"gopkg.in/mgo.v2/bson"
)

var (
	TypeDeviceAdd = "DEVICE_ADDED"
)

type Device struct {
	Id      bson.ObjectId `json:"_id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Serial  string        `json:"serial" bson:"serial"`
	Managed bool          `json:"managed" bson:"managed"`
	Deleted bool          `json:"deleted" bson:"deleted"`
	Status  DeviceStatus  `json:"status" bson:"status"`
}

func (ir *Device) GetObjectId() bson.ObjectId {
	return ir.Id
}

func (ir *Device) SetObjectId(id bson.ObjectId) {
	ir.Id = id
}
