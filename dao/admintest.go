package dao

import (
	. "github.com/tmluthfiana/tryout/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type TestDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	TestCOLLECTION = "test"
)

func (m *TestDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *TestDAO) FindAll() ([]TestModel, error) {
	var test []TestModel
	err := db.C(TestCOLLECTION).Find(bson.M{}).All(&test)
	return test, err
}

func (m *TestDAO) FindById(id string) (TestModel, error) {
	var test TestModel
	err := db.C(TestCOLLECTION).FindId(bson.ObjectIdHex(id)).One(&test)
	return test, err
}

func (m *TestDAO) Insert(test TestModel) error {
	err := db.C(TestCOLLECTION).Insert(&test)
	return err
}

func (m *TestDAO) Delete(test TestModel) error {
	err := db.C(TestCOLLECTION).Remove(&test)
	return err
}

func (m *TestDAO) Update(test TestModel) error {
	err := db.C(TestCOLLECTION).UpdateId(test.ID, &test)
	return err
}
