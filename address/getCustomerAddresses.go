package address

import (
	"context"
	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
)

type GetCustomerAddresses struct {
	NextChain AddressChain
}

func (gca GetCustomerAddresses) Next(obj *AddressObj) *AddressObj {
	var aggRes []struct {
		ConciseAddresses []*ConciseAddress `bson:"conciseAddresses"`
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
					"conciseAddresses": bson.M{
						"$map": bson.M{
							"input": "$addresses",
							"in": bson.M{
								"addr": bson.M{
									"$concat": bson.A{
										"$$this.province.name",
										"/",
										"$$this.city.name",
										"/",
										"$$this.postalAddress.address",
										"-",
										bson.M{"$toString": "$$this.postalAddress.plaque"},
										bson.M{
											"$cond": bson.M{
												"if": bson.M{"$eq": bson.A{bson.M{"$type": "$$this.postalAddress.unit"}, "long"}},
												"then": bson.M{
													"$concat": bson.A{
														"-",
														bson.M{"$toString": "$$this.postalAddress.unit"},
													},
												},
												"else": "",
											},
										},
									},
								},
								"id": "$$this._id",
							},
						},
					},
				},
			},
			bson.M{
				"$project": bson.M{
					"_id":              0,
					"conciseAddresses": 1,
				},
			},
		},
	)
	aggDecoding.All(context.TODO(), &aggRes)
	conciseAddresses := aggRes[0].ConciseAddresses
	obj.ConciseAddresses = conciseAddresses
	if gca.NextChain != nil {
		newObj := gca.NextChain.Next(obj)
		return newObj
	}
	return obj
}
