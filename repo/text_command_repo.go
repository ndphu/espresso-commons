package repo

import (
	"errors"
	"github.com/ndphu/espresso-commons"
	"gopkg.in/mgo.v2"
)

type TextCommandRepo struct {
	Session *mgo.Session
}

func (r *TextCommandRepo) GetCollection() *mgo.Collection {
	return r.Session.DB(commons.DBName).C("text_command")
}

func NewTextCommandRepo(s *mgo.Session) *TextCommandRepo {
	s.SetMode(mgo.Eventual, false)
	return &TextCommandRepo{
		Session: s,
	}
}

func (i *TextCommandRepo) IsSessionConnected() (bool, error) {
	if len(i.Session.LiveServers()) > 0 {
		return true, nil
	} else {
		return false, errors.New("DB is not available")
	}
}
