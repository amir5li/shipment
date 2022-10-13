package address

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateCustomerInfo struct{
	NextChan AddressChain
}

func (uci UpdateCustomerInfo) Next(obj *AddressObj) *AddressObj {
	if obj.UpdateCustomerInfo {
		connection.Customer.UpdateOne(
			context.TODO(),
			bson.M{"phone": obj.UserPhone},
			bson.M{
				"$set": bson.M{
					"firstName": obj.AddressInput.CustomerFirstName,
					"lastName": obj.AddressInput.CustomerLastName,
					"nationalCode": obj.AddressInput.CustomerNationalCode,
				},
			},
		)
	}
	if uci.NextChan != nil {
		newObj := uci.NextChan.Next(obj)
		return newObj
	}
	return obj
}