package Handlers

import (
	"BackendAPI-Instagram/DatabaseConnector"
	"encoding/json"
	"fmt"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	client, ctx := DatabaseConnector.ConnectDatabase()
	defer client.Disconnect(ctx)
	users := DatabaseConnector.GetUsers(client, ctx)
	fmt.Println("This is running")
	data, err := json.Marshal(users)

	if err != nil {
		internalServerError(w, r)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))

}
