package address

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddNewAddress struct {
	NextChain AddressChain
}

func (ana AddNewAddress) Next(obj *AddressObj) *AddressObj {
	var cityDoc models.City
	var provinceDoc models.Province
	findRes := connection.City.FindOne(context.TODO(),bson.M{"_id": obj.AddressInput.AddressCity})
	findRes.Decode(&cityDoc)
	findRes = connection.Province.FindOne(context.TODO(), bson.M{"_id": obj.AddressInput.AddressProvince})
	findRes.Decode(&provinceDoc)
	if obj.AddNewAddress {
		newAddrCity := models.AddressCityAndProvince{
			ID: obj.AddressInput.AddressCity,
			Name: cityDoc.Name,
		}
		newAddrProvince := models.AddressCityAndProvince{
			ID: obj.AddressInput.AddressProvince,
			Name: provinceDoc.Name,
		}
		newAddrPostalAddress := models.PostalAddress{
			Address: obj.AddressInput.AddressPostalAddress,
			PostalCode: obj.AddressInput.AddressPostalCode,
			Plaque: obj.AddressInput.AddressPlaque,
		}
		if obj.AddressInput.AddressUnit != 0 {
			newAddrPostalAddress.Unit = obj.AddressInput.AddressUnit
		}
		var newAddrConsignee models.AddressConsignee
		newAddrConsignee.IsCustomer = obj.AddressInput.ConsigneeIsCustomer
		if obj.AddressInput.ConsigneeIsCustomer {
			newAddrConsignee.FirstName = obj.CustomerInfo.FirstName
			newAddrConsignee.LastName = obj.CustomerInfo.LastName
			newAddrConsignee.NationalCode = obj.CustomerInfo.NationalCode
			newAddrConsignee.Phone = obj.CustomerInfo.Phone
		}else{
			newAddrConsignee.FirstName = obj.AddressInput.ConsigneeFirstName
			newAddrConsignee.LastName = obj.AddressInput.ConsigneeLastName
			newAddrConsignee.NationalCode = obj.AddressInput.ConsigneeNationalCode
			newAddrConsignee.Phone = obj.AddressInput.ConsigneePhone
		}
		newAddrID := primitive.NewObjectID()
		newAddr := models.Address{
			ID: newAddrID,
			City: newAddrCity,
			Province: newAddrProvince,
			PostalAddress: newAddrPostalAddress,
			Consignee: newAddrConsignee,
		}
		connection.Customer.UpdateOne(
			context.TODO(),
			bson.M{"phone": obj.UserPhone},
			bson.M{
				"$push": bson.M{
					"addresses": newAddr,
				},
			},
		)
		obj.InsertedAddressID = newAddrID
	}
	if ana.NextChain != nil {
		newObj := ana.NextChain.Next(obj)
		return newObj
	}
	return obj
}