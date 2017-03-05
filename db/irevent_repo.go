package db

import (
	"gopkg.in/mgo.v2"
)

type IREventRepo struct {
	Session  *mgo.Session
	Database *mgo.Database
}

func (r *IREventRepo) GetCollection() *mgo.Collection {
	return r.Database.C("ir_event")
}
