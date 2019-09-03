package controllers

import (
	// "fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/tmluthfiana/tryout/dao"
	. "github.com/tmluthfiana/tryout/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	// "time"
)

var questiondao = QuestionDAO{}

// GET list of question
func AllQuestion(w http.ResponseWriter, r *http.Request) {
	question, err := questiondao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, question)
}

// GET a question by its ID
func FindQuestion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	question, err := questiondao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid question ID")
		return
	}
	respondWithJson(w, http.StatusOK, question)
}

// POST a new question
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var question QuestionModel
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	question.ID = bson.NewObjectId()
	if err := questiondao.Insert(question); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, question)
}

// PUT update an existing question
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var question QuestionModel
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := questiondao.Update(question); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing question
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var question QuestionModel
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := questiondao.Delete(question); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
