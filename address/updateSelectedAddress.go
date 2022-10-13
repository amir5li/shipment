package address

import (
	"context"
	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateSelectedAddress struct {
	NextChain AddressChain
}

func (usa UpdateSelectedAddress) Next(obj *AddressObj) *AddressObj {
	if obj.UpdateSelectedAddress {
		var cityDoc models.City
		var provinceDoc models.Province
		findRes := connection.City.FindOne(context.TODO(), bson.M{"_id": obj.AddressInput.AddressCity})
		findRes.Decode(&cityDoc)
		findRes = connection.Province.FindOne(context.TODO(), bson.M{"_id": obj.AddressInput.AddressProvince})
		findRes.Decode(&provinceDoc)
		updatingCity := models.AddressCityAndProvince{
			ID: obj.AddressInput.AddressCity,
			Name: cityDoc.Name,
		}
		updatingProvince := models.AddressCityAndProvince{
			ID: obj.AddressInput.AddressProvince,
			Name: provinceDoc.Name,
		}
		updatingConsignee := models.AddressConsignee{}
		if obj.AddressInput.ConsigneeIsCustomer {
			updatingConsignee.FirstName = obj.CustomerInfo.FirstName
			updatingConsignee.LastName = obj.CustomerInfo.LastName
			updatingConsignee.IsCustomer = true
			updatingConsignee.NationalCode = obj.CustomerInfo.NationalCode
			updatingConsignee.Phone = obj.CustomerInfo.Phone
		}else{
			updatingConsignee.IsCustomer = obj.AddressInput.ConsigneeIsCustomer
			updatingConsignee.FirstName = obj.AddressInput.ConsigneeFirstName
			updatingConsignee.LastName = obj.AddressInput.ConsigneeLastName
			updatingConsignee.NationalCode = obj.AddressInput.ConsigneeNationalCode
			updatingConsignee.Phone = obj.AddressInput.ConsigneePhone
		}
		updatingPostalAddress := models.PostalAddress{
			Address: obj.AddressInput.AddressPostalAddress,
			PostalCode: obj.AddressInput.AddressPostalCode,
			Plaque: obj.AddressInput.AddressPlaque,
		}
		if obj.AddressInput.AddressUnit != 0{
			updatingPostalAddress.Unit = obj.AddressInput.AddressUnit
		}
		connection.Customer.UpdateOne(
			context.TODO(),
			bson.M{"phone": obj.UserPhone},
			bson.M{
				"$set": bson.M{
					"addresses.$[targetAddr].city": updatingCity,
					"addresses.$[targetAddr].province": updatingProvince,
					"addresses.$[targetAddr].postalAddress": updatingPostalAddress,
					"addresses.$[targetAddr].consignee": updatingConsignee,
				},
			},
		)
	}
	if usa.NextChain != nil {
		newObj := usa.NextChain.Next(obj)
		return newObj
	}
	return obj
}