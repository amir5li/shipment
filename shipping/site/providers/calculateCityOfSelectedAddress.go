package providers

import (
	"context"
	"fmt"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func calculateCityOfSelectedAddress(ctx context.Context, customerID primitive.ObjectID, addressID primitive.ObjectID) primitive.ObjectID {
	var aggRes []struct{
		CityID primitive.ObjectID `bson:"cityID"`
	}
	aggDecoding, err := connection.Customer.Aggregate(
		ctx, 
		bson.A{
			bson.M{
				"$match": bson.M{
					"_id": customerID,
				},
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
	fmt.Println(err)
	aggDecoding.All(ctx, &aggRes)
	fmt.Println(err, "city", aggRes)
	result := aggRes[0].CityID
	fmt.Println("city")
	return result
}