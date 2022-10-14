package shippingAdmin

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddPricePlan(ctx context.Context, inp AddPricePlanInput) (*Res, error){
	minWeightErr := _validateMinWeight(inp.MinWeight)
	if minWeightErr != nil {
		return nil, minWeightErr
	}
	maxWeightErr := _validateMaxWeight(inp.MaxWeight)
	if maxWeightErr != nil {
		return nil, maxWeightErr
	}
	priceErr := _validatePrice(inp.Price)
	if priceErr != nil {
		return nil, priceErr
	}
	insertingPricePlan := models.PricePlan{
		ID: primitive.NewObjectID(),
		MaxWeight: inp.MaxWeight,
		MinWeight: inp.MinWeight,
		Price: inp.Price,
	}
	_, err := connection.ShippingMethod.UpdateByID(
		ctx,
		inp.MethodID,
		bson.M{
			"$push": bson.M{
				"pricePlans": bson.M{
					"$each": bson.A{insertingPricePlan},
					"$sort": bson.M{"price": 1},
				},
			},
		},
	)
	if err != nil {
		return nil, DBError
	}
	return &Res{Messages.SuccessMsg, Messages.SuccessCode}, nil
}