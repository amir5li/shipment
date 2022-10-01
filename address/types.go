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

type AddressObj struct {
	Form []*AddressSection `json:"form"`
	Action string `json:"action"`
	SelectedAddressID primitive.ObjectID `json:"selectedAddressID"`
	ShowAddresses []ShowAddressField `json:"showAddresses"`
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