package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BasketItem struct {
	VarID primitive.ObjectID `bson:"varID"`
	Count uint `bson:"count"`
}

type Basket struct {
	ID primitive.ObjectID `bson:"_id"`
	CustomerID primitive.ObjectID `bson:"customerID"`
	Items []BasketItem `bson:"items"`
}