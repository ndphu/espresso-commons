package repo

import (
	"errors"
	"github.com/ndphu/espresso-commons"
	"gopkg.in/mgo.v2"
)

type IREventRepo struct {
	Session *mgo.Session
}

func (r *IREventRepo) GetCollection() *mgo.Collection {
	return r.Session.DB(commons.DBName).C("ir_event")
}

func NewIREventRepo(s *mgo.Session) *IREventRepo {
	s.SetMode(mgo.Eventual, false)
	return &IREventRepo{
		Session: s,
	}
}

func (i *IREventRepo) IsSessionConnected() (bool, error) {
	if len(i.Session.LiveServers()) > 0 {
		return true, nil
	} else {
		return false, errors.New("DB is not available")
	}
}
