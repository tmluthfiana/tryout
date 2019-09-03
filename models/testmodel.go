package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TestModel struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	NameTest    string        `bson:"name_test" json:"name_test"`
	DateTest    time.Time     `bson:"date_test" json:"date_test"`
	Description string        `bson:"description" json:"description"`
	Category    string        `bson:"category" json:"category"`
}
