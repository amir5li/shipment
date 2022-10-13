package address

import (
	"context"
	"fmt"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SetNewAddressDefaltIsNeeded struct {
	NextChain AddressChain
}

func (snad SetNewAddressDefaltIsNeeded) Next(obj *AddressObj)*AddressObj {
	var aggRes []struct {
		AddressCount int `bson:"addressCount"`
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
					"addressCount": bson.M{
						"$size": "$addresses",
					},
				},
			},
		},
	)
	aggDecoding.All(context.TODO(), &aggRes)
	fmt.Println("addressCount", aggRes[0].AddressCount)
	if aggRes[0].AddressCount == 1 {
		connection.Customer.UpdateOne(
			context.TODO(),
			bson.M{"phone": obj.UserPhone},
			bson.M{
				"$set": bson.M{
					"addresses.$[targetAddr].isDefault": true,
				},
			},
			options.Update().SetArrayFilters(options.ArrayFilters{
				Filters: bson.A{bson.M{"targetAddr._id": obj.InsertedAddressID}},
			}),
		)
		obj.SessionAddressID = obj.InsertedAddressID
	}
	if snad.NextChain != nil {
		newObj := snad.NextChain.Next(obj)
		return newObj
	}
	return obj
}