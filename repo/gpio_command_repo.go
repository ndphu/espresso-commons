package repo

import (
	"errors"
	"github.com/ndphu/espresso-commons"
	"gopkg.in/mgo.v2"
)

type GPIOCommandRepo struct {
	Session *mgo.Session
}

func (r *GPIOCommandRepo) GetCollection() *mgo.Collection {
	return r.Session.DB(commons.DBName).C("gpio_command")
}

func NewGPIOCommandRepo(s *mgo.Session) *GPIOCommandRepo {
	s.SetMode(mgo.Eventual, false)
	return &GPIOCommandRepo{
		Session: s,
	}
}

func (i *GPIOCommandRepo) IsSessionConnected() (bool, error) {
	if len(i.Session.LiveServers()) > 0 {
		return true, nil
	} else {
		return false, errors.New("DB is not available")
	}
}
