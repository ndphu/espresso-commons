package repo

import (
	"errors"
	"github.com/ndphu/espresso-commons"
	"gopkg.in/mgo.v2"
)

type DeviceRepo struct {
	Session *mgo.Session
}

func (r *DeviceRepo) GetCollection() *mgo.Collection {
	return r.Session.DB(commons.DBName).C("device")
}

func NewDeviceRepo(s *mgo.Session) *DeviceRepo {
	s.SetMode(mgo.Eventual, false)
	return &DeviceRepo{
		Session: s,
	}
}

func (i *DeviceRepo) IsSessionConnected() (bool, error) {
	if len(i.Session.LiveServers()) > 0 {
		return true, nil
	} else {
		return false, errors.New("DB is not available")
	}
}
