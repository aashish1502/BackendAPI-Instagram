package DatabaseConnector

import (
	"BackendAPI-Instagram/Structures"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func GetUsers(client *mongo.Client, ctx context.Context) []Structures.User {

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	queryDatabase := client.Database("Users")
	queryCollection := queryDatabase.Collection("UsersTest")

	cursor, err := queryCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	var users []Structures.User

	for cursor.Next(ctx) {
		var user = Structures.User{}
		if cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err = cursor.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
	return users
}
