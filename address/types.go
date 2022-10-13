package address

import (
	"github.com/amir5li/shipment/location"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressField struct {
	Name string `json:"name"`
	Label string `json:"label"`
	Type string `json:"type"`
	Value interface{} `json:"value"`
	Required bool `json:"required"`
	Disabled bool `json:"disabled"`
	Error *string `json:"error"`
	Options []location.Loc `json:"options"`
}
type AddressSection struct {
	Title string `json:"title"`
	Fields []*AddressField `json:"fields"`
}
type ShowAddressField struct {
	Address string `json:"address"`
	ID primitive.ObjectID `json:"id"`
}

type CustomerInfo struct {
	FirstName string
	LastName string
	Phone string
	NationalCode string
}

type AddressObj struct {
	Form []*AddressSection `json:"form"`
	ConciseAddresses []ConciseAddress `json:"addresses"`
	SelectedAddressID primitive.ObjectID `json:"selectedAddressID"`
	InsertedAddressID primitive.ObjectID 
	SessionAddressID primitive.ObjectID 
	ShowAddresses []ShowAddressField `json:"showAddresses"`
	AddressInput AddressInput
	CustomerInfo CustomerInfo
	NeedUpdateCustomerInfo bool
	UpdateCustomerInfo bool
	UpdateSelectedAddress bool
	AddNewAddress bool
	UserPhone string
}

type AddressInput struct {
	CustomerFirstName string `json:"customerFirstName"`
	CustomerLastName string `json:"customerLastName"`
	CustomerNationalCode string `json:"customerNationalCode"`
	ConsigneeIsCustomer bool `json:"isCustomer"`
	ConsigneeFirstName string `json:"consigneeFirstName"`
	ConsigneeLastName string `json:"consigneeLastName"`
	ConsigneeNationalCode string `json:"consigneeNationalCode"`
	ConsigneePhone string `json:"consigneePhone"`
	AddressCity primitive.ObjectID `json:"city"`
	AddressProvince primitive.ObjectID `json:"province"`
	AddressPostalAddress string `json:"postalAddress"`
	AddressPostalCode string `json:"postalCode"`
	AddressUnit uint `json:"unit"`
	AddressPlaque uint `json:"plaque"`
}

type ConciseAddress struct {
	Address string `bson:"addr" json:"address"`
	ID primitive.ObjectID `bson:"id" json:"id"`
}

type EditAddressInput struct {
	SelectedAddress primitive.ObjectID `json:"selectedAddress"`
}

type SubmitEditAddressInput struct {
	SelectedAddressID primitive.ObjectID `json:"selectedAddress"`
	AddressInput AddressInput `json:"addressInput"`
}