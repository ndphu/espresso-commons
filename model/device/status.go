package device

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type DeviceStatus struct {
	Id        bson.ObjectId `json:"_id" bson:"_id"`
	Serial    string        `json:"serial" bson:"serial"`
	Online    bool          `json:"online" bson:"online"`
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
	Free      int           `json:"free" bson:"free"`
	Uptime    int           `json:"uptime" bson:"uptime"`
}

func (d *DeviceStatus) GetObjectId() bson.ObjectId {
	return d.Id
}

func (d *DeviceStatus) SetObjectId(id bson.ObjectId) {
	d.Id = id
}
