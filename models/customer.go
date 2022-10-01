package models

type Customer struct {
	FirstName string `bson:"firstName"`
	LastName string `bson:"lastName"`
	Phone string `bson:"phone"`
	NationalCode string `bson:"nationalCode"`
	Addresses []Address
}