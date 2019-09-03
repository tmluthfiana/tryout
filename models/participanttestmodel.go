package models

import (
	"gopkg.in/mgo.v2/bson"
)

type ParticipantTestModel struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	NameTest    string        `bson:"name_test" json:"name_test"`
	IdTest      bson.ObjectId `bson:"id_test" json:"id_test"`
	Participant []ParticipantModel
}

type ParticipantModel struct {
	StudentName  string        `bson:"student_name" json:"student_name"`
	StudentId    bson.ObjectId `bson:"student_id" json:"student_id"`
	StudentEmail string        `bson:"student_email" json:"student_email"`
}
