package Handlers

import (
	"BackendAPI-Instagram/DatabaseConnector"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func getByUser(w http.ResponseWriter, r *http.Request) {
	client, ctx := DatabaseConnector.ConnectDatabase()
	defer client.Disconnect(ctx)

	matches := getpostbyUserRe.FindStringSubmatch(r.URL.Path)
	id, _ := primitive.ObjectIDFromHex(matches[1])

	user := DatabaseConnector.GetUserByID(client, ctx, id)

	mail := (user[0].Email)

	posts := DatabaseConnector.FetchByEmail(client, ctx, mail)

	if posts == nil {
		notFound(w, r)
	} else {
		data, _ := json.Marshal(posts)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

}
