package repo

import (
	"errors"
	"github.com/ndphu/espresso-commons"
	"gopkg.in/mgo.v2"
)

type DeviceStatusRepo struct {
	Session *mgo.Session
}

func (r *DeviceStatusRepo) GetCollection() *mgo.Collection {
	return r.Session.DB(commons.DBName).C("device_status")
}

func NewDeviceStatusRepo(s *mgo.Session) *DeviceStatusRepo {
	s.SetMode(mgo.Eventual, false)
	return &DeviceStatusRepo{
		Session: s,
	}
}

func (i *DeviceStatusRepo) IsSessionConnected() (bool, error) {
	if len(i.Session.LiveServers()) > 0 {
		return true, nil
	} else {
		return false, errors.New("DB is not available")
	}
}
