package models

import (
	"gopkg.in/mgo.v2/bson"
)

type QuestionModel struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	NameTest string        `bson:"name_test" json:"name_test"`
	IdTest   bson.ObjectId `bson:"id_test" json:"id_test"`
	Question string        `bson:"question" json:"question"`
	Answer   string        `bson:"answer" json:"answer"`
	OptionA  string        `bson:"option_a" json:"option_a"`
	OptionB  string        `bson:"option_b" json:"option_b"`
	OptionC  string        `bson:"option_c" json:"option_c"`
	OptionD  string        `bson:"option_d" json:"option_d"`
	OptionE  string        `bson:"option_e" json:"option_e"`
}
