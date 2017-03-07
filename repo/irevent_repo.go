package repo

import (
	"github.com/ndphu/espresso-commons"
	"gopkg.in/mgo.v2"
)

type IREventRepo struct {
	Session *mgo.Session
}

func (r *IREventRepo) GetCollection() *mgo.Collection {
	return r.Session.DB(commons.DBName).C("ir_event")
}
