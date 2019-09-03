package controllers

import (
	// "fmt"
	// "encoding/json"
	. "github.com/tmluthfiana/tryout/dao"
	"net/http"
)

var ptestdao = ParticipantTestDAO{}

// GET list of ptestdao
func AllParticipantTest(w http.ResponseWriter, r *http.Request) {
	ptestdao, err := ptestdao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, ptestdao)
}
