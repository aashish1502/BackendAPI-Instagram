package Handlers

import (
	"BackendAPI-Instagram/DatabaseConnector"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GETByID(w http.ResponseWriter, r *http.Request) {

	cliet, ctx := DatabaseConnector.ConnectDatabase()
	defer cliet.Disconnect(ctx)
	//fmt.Println("This is running")
	matches := getUserRe.FindStringSubmatch(r.URL.Path)
	id, _ := primitive.ObjectIDFromHex(matches[1])
	user := DatabaseConnector.GetUserByID(cliet, ctx, id)
	if user == nil {
		notFound(w, r)
	} else {
		data, err := json.Marshal(user)
		if err != nil {
			internalServerError(w, r)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
