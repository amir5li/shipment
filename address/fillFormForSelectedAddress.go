package address

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson"
)

type FillFormForSelectedAddress struct {
	NextChain AddressChain
}

func (ffsa FillFormForSelectedAddress) Next(obj *AddressObj) *AddressObj {
	var aggRes []struct {
		SelectedAddress models.Address `bson:"selectedAddress"`
	}
	aggDecoding, _ := connection.Customer.Aggregate(
		context.TODO(),
		bson.A{
			bson.M{
				"$match": bson.M{
					"phone": obj.UserPhone,
				},
			},
			bson.M{
				"$set": bson.M{
					"selectedAddress": bson.M{
						"$reduce": bson.M{
							"input": "$addresses",
							"initialValue": models.Address{},
							"in": bson.M{
								"$cond": bson.M{
									"if": bson.M{
										"$eq": bson.A{
											"$$this._id",
											obj.SelectedAddressID,
										},
									},
									"then": "$$this",
									"else": "$$value",
								},
							},
						},
					},
				},
			},
			bson.M{
				"$project": bson.M{
					"_id": 0,
					"selectedAddress": 1,
				},
			},
		},
	)
	aggDecoding.All(context.TODO(), &aggRes)
	selectedAddr := aggRes[0].SelectedAddress
	for _, sec := range obj.Form {
		for _, field := range sec.Fields {
			switch field.Name {
			case ConsigneeIsCustomerFieldName:
				field.Value = selectedAddr.Consignee.IsCustomer
			case ConsigneeFirstNameFieldName:
				field.Value = selectedAddr.Consignee.FirstName
			case ConsigneeLastNameFieldName:
				field.Value = selectedAddr.Consignee.LastName
			case ConsigneeNationalCodeFieldName:
				field.Value = selectedAddr.Consignee.NationalCode
			case ConsigneePhoneFieldName:
				field.Value = selectedAddr.Consignee.Phone
			case AddressCityFieldName:
				field.Value = selectedAddr.City.ID
			case AddressProvinceFieldName:
				field.Value = selectedAddr.Province.ID
			case AddressPostalAddressFieldName:
				field.Value = selectedAddr.PostalAddress.Address
			case AddressPostalCodeFieldName:
				field.Value = selectedAddr.PostalAddress.PostalCode
			case AddressPlaqueFieldName:
				field.Value = selectedAddr.PostalAddress.Plaque
			case AddressUnitFieldName:
				field.Value = selectedAddr.PostalAddress.Unit
			}
		}
	}
	if ffsa.NextChain != nil {
		newObj := ffsa.NextChain.Next(obj)
		return newObj
	}
	return obj
}