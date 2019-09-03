package controllers

import (
	// "fmt"
	// "encoding/json"
	"github.com/gorilla/mux"
	// . "github.com/tmluthfiana/tryout/dao"
	// . "github.com/tmluthfiana/tryout/models"
	// "gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func StartTheTest(w http.ResponseWriter, r *http.Request) {
	timer1 := time.NewTimer(15 * time.Minute)
	<-timer1.C
	w.Write([]byte("Time is Up"))
}

func GetTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	question, err := questiondao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid question ID")
		return
	}
	respondWithJson(w, http.StatusOK, question)
}
