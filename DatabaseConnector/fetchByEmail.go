package DatabaseConnector

import (
	"BackendAPI-Instagram/Structures"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func FetchByEmail(client *mongo.Client, ctx context.Context, email string) []Structures.Post {

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	queryDatabase := client.Database("Users")
	queryCollection := queryDatabase.Collection("Posts")
	cursor, err := queryCollection.Find(ctx, bson.M{"email": email})
	if err != nil {
		log.Fatal(err)
	}

	var posts []Structures.Post

	for cursor.Next(ctx) {
		var post = Structures.Post{}
		if cursor.Decode(&post); err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}

	return posts

}
