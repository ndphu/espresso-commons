package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

var (
	TypeIREventAdd = "IR_EVENT_ADDED"
)

type IRMessage struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	Button        string        `json:"button"`
	RemoteName    string        `json:"remoteName"`
	Code          uint64        `json:"code"`
	Repeat        int64         `json:"repeat"`
	Source        string        `json:"source"`
	Timestamp     time.Time     `json:"timestamp"`
	UnixTimestamp int64         `json:"unixTimestamp" bson:"-"`
}

func (ir *IRMessage) GetObjectId() bson.ObjectId {
	return ir.Id
}

func (ir *IRMessage) SetObjectId(id bson.ObjectId) {
	ir.Id = id
}
