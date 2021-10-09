package Structures

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}
