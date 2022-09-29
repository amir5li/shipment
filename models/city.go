package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type City struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	ProvinceID primitive.ObjectID `bson:"provinceID" json:"provinceID"`
}