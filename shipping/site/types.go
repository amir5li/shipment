package shippingSite

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetShippingDataInput struct {
	MethodID primitive.ObjectID `json:"methodID"`
}