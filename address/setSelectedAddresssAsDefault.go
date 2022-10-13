package address

import (
	"context"
	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SetSelectedAddressAsDefault struct {
	NextChain AddressChain
}

func (ssad SetSelectedAddressAsDefault) Next(obj *AddressObj) *AddressObj {
	connection.Customer.UpdateOne(
		context.TODO(),
		bson.M{"phone": obj.UserPhone},
		bson.M{
			"$set": bson.M{
				"addresses.$[targetAddr].isDefault": true,
				"addresses.$[otherAddrs].isDefault": false,
			},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: bson.A{
				bson.M{
					"targetAddr._id": obj.SelectedAddressID,
				},
				bson.M{
					"otherAddrs._id": bson.M{"$ne": obj.SelectedAddressID},
				},
			},
		}),
	)
	if ssad.NextChain != nil {
		newObj := ssad.NextChain.Next(obj)
		return newObj
	}
	return obj
}
