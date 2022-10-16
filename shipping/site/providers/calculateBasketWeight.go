package providers

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CalculateBasketWeight(ctx context.Context, customerID primitive.ObjectID) uint {
	var aggRes []struct{
		TotalWeight uint `bson:"totalWeight"`
	}
	aggDecoding, _ := connection.Basket.Aggregate(
		ctx,
		bson.A{
			bson.M{
				"customerID": customerID,
			},
			bson.M{
				"$lookup": bson.M{
					"from": "products",
					"localField": "items.varID",
					"foreignField": "variations._id",
					"as": "prdDoc",
				},
			},
			bson.M{
				"$set": bson.M{
					"totalWeight": bson.M{
						"$reduce": bson.M{
							"input": "$prdDoc",
							"initialValue": 0,
							"in": bson.M{
								"$sum": bson.A{"$$value", "$$this.weight"},
							},
						},
					},
				},
			},
			bson.M{
				"$project": bson.M{
					"_id": 0,
					"totalWeight": 1,
				},
			},
		},
		
	)
	aggDecoding.All(ctx, &aggRes)
	result := aggRes[0].TotalWeight
	return result
}