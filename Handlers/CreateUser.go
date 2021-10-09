package Handlers

import (
	"BackendAPI-Instagram/DatabaseConnector"
	"BackendAPI-Instagram/Structures"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var u Structures.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		internalServerError(w, r)
	}

	//fmt.Println(u)
	DatabaseConnector.InsertUser(u)

}
