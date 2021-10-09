package Handlers

import (
	"BackendAPI-Instagram/DatabaseConnector"
	"BackendAPI-Instagram/Structures"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {

	var p Structures.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	fmt.Println(p)
	if err != nil {
		internalServerError(w, r)
		return
	}
	DatabaseConnector.InsertPost(p)

}
