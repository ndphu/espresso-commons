package dao

import (
	"errors"
	"github.com/ndphu/manga-crawler/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrorMessageDBDisonnected string = "connection to mongodb is not available"
)

type Repo interface {
	GetCollection() *mgo.Collection
	IsSessionConnected() (bool, error)
}

func Count(r Repo, selector bson.M) (int, error) {
	if _, err := r.IsSessionConnected(); err != nil {
		return -1, err
	}
	return r.GetCollection().Count()
}

func CountBy(r Repo, selector bson.M) (int, error) {
	if _, err := r.IsSessionConnected(); err != nil {
		return -1, err
	}
	return r.GetCollection().Find(selector).Count()
}

func FindOne(r Repo, selector bson.M, result interface{}) error {
	if _, err := r.IsSessionConnected(); err != nil {
		return err
	}
	return r.GetCollection().Find(selector).One(result)
}

func FindAll(r Repo, selector bson.M, skip int, limit int, result interface{}) error {
	if _, err := r.IsSessionConnected(); err != nil {
		return err
	}
	return r.GetCollection().Find(selector).Skip(skip).Limit(limit).All(result)
}

func FindAllWithSort(r Repo, selector bson.M, skip int, limit int, sort string, result interface{}) error {
	if _, err := r.IsSessionConnected(); err != nil {
		return err
	}
	return r.GetCollection().Find(selector).Sort(sort).Skip(skip).Limit(limit).All(result)
}

func FindById(r Repo, id bson.ObjectId, result interface{}) error {
	if _, err := r.IsSessionConnected(); err != nil {
		return err
	}
	return FindOne(r, bson.M{"_id": id}, result)
}

func Insert(r Repo, m model.Model) error {
	if _, err := r.IsSessionConnected(); err != nil {
		return err
	}
	m.SetObjectId(bson.NewObjectId())
	return r.GetCollection().Insert(m)
}

func InsertAll(r Repo, models []interface{}) error {
	if _, err := r.IsSessionConnected(); err != nil {
		return err
	}
	for i := 0; i < len(models); i++ {
		models[i].(model.Model).SetObjectId(bson.NewObjectId())
	}
	return r.GetCollection().Insert(models...)
}

func Update(r Repo, m model.Model) error {
	if _, err := r.IsSessionConnected(); err != nil {
		return err
	}
	if len(m.GetObjectId().Hex()) == 0 {
		return errors.New("ObjectId is required for update")
	}
	return r.GetCollection().UpdateId(m.GetObjectId(), m)
}

func Delete(r Repo, id bson.ObjectId) error {
	if _, err := r.IsSessionConnected(); err != nil {
		return err
	}
	return r.GetCollection().RemoveId(id)
}

func DeleteAll(r Repo) error {
	if _, err := r.IsSessionConnected(); err != nil {
		return err
	}
	return DeleteAllByCondition(r, bson.M{})
}

func DeleteAllByCondition(r Repo, selector bson.M) error {
	if _, err := r.IsSessionConnected(); err != nil {
		return err
	}
	_, err := r.GetCollection().RemoveAll(selector)
	return err
}
