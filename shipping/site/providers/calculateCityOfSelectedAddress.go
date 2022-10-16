package providers

import (
	"context"
	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CalculateCityOfSelectedAddress(ctx context.Context, customerID primitive.ObjectID, addressID primitive.ObjectID) primitive.ObjectID {
	var aggRes []struct{
		CityID primitive.ObjectID `bson:"cityID"`
	}
	aggDecoding, _ := connection.Customer.Aggregate(
		ctx, 
		bson.A{
			bson.M{
				"$match": customerID,
			},
			bson.M{
				"$set": bson.M{
					"cityID": bson.M{
						"$reduce": bson.M{
							"input": "$addresses",
							"initialValue": nil,
							"in": bson.M{
								"$cond": bson.M{
									"if": bson.M{
										"$and": bson.A{
											bson.M{
												"$eq": bson.A{"$$value", nil},
											},
											bson.M{
												"$eq": bson.A{"$$this._id", addressID},
											},
										},
									},
									"then": "$$this.city.id",
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
					"cityID": 1,
				},
			},
		},
	)
	aggDecoding.All(ctx, &aggRes)
	result := aggRes[0].CityID
	return result
}