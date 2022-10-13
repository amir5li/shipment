package address

import (
	"context"
	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FindSelectedAddress struct {
	NextChain AddressChain
}

func (fsa FindSelectedAddress) Next(obj *AddressObj) *AddressObj {
	if len(obj.ConciseAddresses) == 0 {
		return obj
	}
	var targetAddr primitive.ObjectID
	if !obj.SelectedAddressID.IsZero() {
		targetAddr = obj.SelectedAddressID
	} else {
		var defaultAddr []struct {
			DefaultAddress primitive.ObjectID `bson:"defaultAddress"`
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
						"defaultAddress": bson.M{
							"$reduce": bson.M{
								"input":        "$addresses",
								"initialValue": nil,
								"in": bson.M{
									"if": bson.M{
										"$and": bson.A{
											bson.M{
												"$eq": bson.A{"$$value", nil},
											},
											bson.M{"$eq": bson.A{"$$this.isDefault", true}},
										},
									},
								},
							},
						},
					},
				},
				bson.M{
					"$project": bson.M{
						"_id":            0,
						"defaultAddress": 1,
					},
				},
			},
		)
		aggDecoding.All(context.TODO(), &defaultAddr)
		targetAddr = defaultAddr[0].DefaultAddress
	}
	for _, addr := range obj.ConciseAddresses {
		if addr.ID.Hex() == targetAddr.Hex() {
			addr.Selected = true
		}
	}
	obj.SessionAddressID = targetAddr
	if fsa.NextChain != nil {
		newObj := fsa.NextChain.Next(obj)
		return newObj
	}
	return obj
}
