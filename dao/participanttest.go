package dao

import (
	. "github.com/tmluthfiana/tryout/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ParticipantTestDAO struct {
	Server   string
	Database string
}

const (
	pTestCOLLECTION = "participanttest"
)

func (m *ParticipantTestDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ParticipantTestDAO) FindAll() ([]ParticipantTestModel, error) {
	var ptest []ParticipantTestModel
	err := db.C(pTestCOLLECTION).Find(bson.M{}).All(&ptest)
	return ptest, err
}
