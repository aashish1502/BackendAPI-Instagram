package DatabaseConnector

import (
	"BackendAPI-Instagram/Structures"
	"log"
)

func InsertPost(post Structures.Post) {

	client, ctx := ConnectDatabase()
	defer client.Disconnect(ctx)

	queryDatabase := client.Database("Users")
	queryCollection := queryDatabase.Collection("Posts")

	if UserExists(post) == true {
		_, err := queryCollection.InsertOne(ctx, post)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
