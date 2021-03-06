package Structures

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

type Post struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email   string             `json:"email,omitempty" bson:"email,omitempty"`
	Caption string             `json:"caption,omitempty"bson:"caption,omitempty"`
	URL     string             `json:"url,omitempty" bson:"url,omitempty"`
	Time    primitive.DateTime `json:"time,omitempty" bson:"time,omitempty"`
}
