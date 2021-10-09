package DatabaseConnector

import (
	"BackendAPI-Instagram/Structures"
	"log"
)

func InsertUser(user Structures.User) {

	client, ctx := ConnectDatabase()
	defer client.Disconnect(ctx)

	queryDatabase := client.Database("Users")
	queryCollection := queryDatabase.Collection("UsersTest")

	_, err := queryCollection.InsertOne(ctx, user)

	if err != nil {
		log.Fatalln(err)
	}

}
