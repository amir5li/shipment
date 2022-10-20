package providers

import (
	"context"
	"github.com/amir5li/shipment/models"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func calculateBasketWeight(ctx context.Context, customerID primitive.ObjectID) uint {
	var variationsWeight []struct {
		VarID  primitive.ObjectID `bson:"varID"`
		Weight uint               `bson:"weight"`
	}
	var targetBasket models.Basket
	findRes := connection.Basket.FindOne(
		ctx,
		bson.M{"customerID": customerID},
	)
	findRes.Decode(&targetBasket)
	var varIDs []primitive.ObjectID
	for _, variation := range targetBasket.Items {
		varIDs = append(varIDs, variation.VarID)
	}
	aggDecoding, _ := connection.Product.Aggregate(
		ctx,
		bson.A{
			bson.M{
				"variations._id": bson.M{
					"$in": varIDs,
				},
			},
			bson.M{
				"$set": bson.M{
					"varID": bson.M{
						"$reduce": bson.M{
							"input":        "$variations",
							"initialValue": nil,
							"in": bson.M{
								"$cond": bson.M{
									"if": bson.M{
										"$and": bson.A{
											bson.M{
												"$eq": bson.A{"$$value", nil},
											},
											bson.M{
												"$in": bson.A{"$$this._id", varIDs},
											},
										},
									},
									"then": "$$this._id",
									"else": "$$value",
								},
							},
						},
					},
				},
			},
			bson.M{
				"$project": bson.M{
					"_id":    0,
					"weight": 1,
					"varID":  1,
				},
			},
		},
	)
	aggDecoding.All(ctx, &variationsWeight)
	var totalWeight uint
	for _, basketItem := range targetBasket.Items {
		for _, varW := range variationsWeight {
			if varW.VarID.Hex() == basketItem.VarID.Hex() {
				totalWeight += basketItem.Count + varW.Weight
			}
		}
	}
	return totalWeight
}