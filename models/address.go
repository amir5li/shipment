package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressCityAndProvince struct {
	ID primitive.ObjectID `bson:"id"`
	Name string `bson:"name"`
}
type AddressConsignee struct {
	IsCustomer bool `bson:"isCustomer"`
	FirstName string `bson:"firstName"`
	LastName string `bson:"lastName"`
	NationalCode string `bson:"nationalCode"`
	Phone string `bson:"phone"`
}
type PostalAddress struct {
	Plaque uint `bson:"plaque"`
	Unit uint `bson:"unit,omitempty"`
	PostalCode string `bson:"postalCode"`
	Address string `bson:"address"`
}
type Address struct {
	ID primitive.ObjectID `bson:"_id"`
	Province AddressCityAndProvince `bson:"province"`	
	City AddressCityAndProvince `bson:"city"`
	PostalAddress PostalAddress  `bson:"postalAddress"`
	Consignee AddressConsignee `bson:"consignee"`
	IsDefault bool `bson:"isDefault"`
}