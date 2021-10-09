package main

import (
	"BackendAPI-Instagram/DatabaseConnector"
	"BackendAPI-Instagram/Handlers"
	"fmt"
	"net/http"
)

func main() {

	client, ctx := DatabaseConnector.ConnectDatabase()
	defer client.Disconnect(ctx)
	fmt.Println("Starting Server....")

	http.HandleFunc("/users", Handlers.ServeHTTP)
	http.HandleFunc("/users/", Handlers.ServeHTTP)
	http.ListenAndServe("localhost:8080", nil)

}
