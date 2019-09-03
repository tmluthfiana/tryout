package main

import (
	// "fmt"
	// "encoding/json"
	"github.com/gorilla/mux"
	. "github.com/tmluthfiana/tryout/config"
	. "github.com/tmluthfiana/tryout/controllers"
	. "github.com/tmluthfiana/tryout/dao"
	"log"
	"net/http"
)

var config = Config{}
var test = TestDAO{}
var question = QuestionDAO{}
var ptest = ParticipantTestDAO{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	test.Server = config.Server
	test.Database = config.Database
	test.Connect()

	question.Server = config.Server
	question.Database = config.Database
	question.Connect()

	ptest.Server = config.Server
	ptest.Database = config.Database
	ptest.Connect()
}

// Define HTTP request routes

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/admintest", AllTest).Methods("GET")
	r.HandleFunc("/admintest", CreateTest).Methods("POST")
	r.HandleFunc("/admintest", UpdateTest).Methods("PUT")
	r.HandleFunc("/admintest", DeleteTest).Methods("DELETE")
	r.HandleFunc("/admintest/{id}", FindTest).Methods("GET")
	r.HandleFunc("/adminquestion", AllQuestion).Methods("GET")
	r.HandleFunc("/adminquestion", CreateQuestion).Methods("POST")
	r.HandleFunc("/adminquestion", UpdateQuestion).Methods("PUT")
	r.HandleFunc("/adminquestion", DeleteQuestion).Methods("DELETE")
	r.HandleFunc("/adminquestion/{id}", FindQuestion).Methods("GET")
	r.HandleFunc("/partisipanttest", AllParticipantTest).Methods("GET")
	r.HandleFunc("/taketest", StartTheTest).Methods("GET")
	r.HandleFunc("/taketest/{id}", GetTest).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
