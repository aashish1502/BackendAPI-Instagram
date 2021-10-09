package DatabaseConnector

import (
	"BackendAPI-Instagram/Structures"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func FetchPosts(client *mongo.Client, ctx context.Context, id string) []Structures.Post {

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	objId, _ := primitive.ObjectIDFromHex(id)
	queryDatabase := client.Database("Users")
	queryCollection := queryDatabase.Collection("Posts")

	cursor, err := queryCollection.Find(ctx, bson.M{"_id": objId})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
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
