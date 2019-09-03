package dao

import (
	. "github.com/tmluthfiana/tryout/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type QuestionDAO struct {
	Server   string
	Database string
}

const (
	QuestionCOLLECTION = "question"
)

func (m *QuestionDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *QuestionDAO) FindAll() ([]QuestionModel, error) {
	var question []QuestionModel
	err := db.C(QuestionCOLLECTION).Find(bson.M{}).All(&question)
	return question, err
}

func (m *QuestionDAO) FindById(id string) (QuestionModel, error) {
	var question QuestionModel
	err := db.C(QuestionCOLLECTION).FindId(bson.ObjectIdHex(id)).One(&question)
	return question, err
}

func (m *QuestionDAO) Insert(question QuestionModel) error {
	err := db.C(QuestionCOLLECTION).Insert(&question)
	return err
}

func (m *QuestionDAO) Delete(question QuestionModel) error {
	err := db.C(QuestionCOLLECTION).Remove(&question)
	return err
}

func (m *QuestionDAO) Update(question QuestionModel) error {
	err := db.C(QuestionCOLLECTION).UpdateId(question.ID, &question)
	return err
}
