package controllers

import (
	// "fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/tmluthfiana/tryout/dao"
	// . "github.com/tmluthfiana/tryout/helper"
	. "github.com/tmluthfiana/tryout/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	// "time"
)

var testdao = TestDAO{}

// GET list of test
func AllTest(w http.ResponseWriter, r *http.Request) {
	test, err := testdao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, test)
}

// GET a test by its ID
func FindTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	test, err := testdao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Test ID")
		return
	}
	respondWithJson(w, http.StatusOK, test)
}

// POST a new test
func CreateTest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var test TestModel
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	test.ID = bson.NewObjectId()
	if err := testdao.Insert(test); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, test)
}

// PUT update an existing test
func UpdateTest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var test TestModel
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := testdao.Update(test); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing test
func DeleteTest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var test TestModel
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := testdao.Delete(test); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
