package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Province struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
}