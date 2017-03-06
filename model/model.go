package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Model interface {
	GetObjectId() bson.ObjectId
	SetObjectId(bson.ObjectId)
}
