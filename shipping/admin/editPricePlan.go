package shippingAdmin

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EditPricePlan(ctx context.Context, inp EditPricePlanInput) (*Res, error){
	maxWeightErr := _validateMaxWeight(inp.MaxWeight)
	if maxWeightErr != nil {
		return nil, maxWeightErr
	}
	minWeightErr := _validateMinWeight(inp.MinWeight)
	if minWeightErr != nil {
		return nil, minWeightErr
	}
	priceErr := _validatePrice(inp.Price)
	if priceErr != nil {
		return nil, priceErr
	}
	_, err := connection.ShippingMethod.UpdateByID(
		ctx, 
		inp.MethodID,
		bson.M{
			"$set": bson.M{
				"pricePlans.$[targetPlan].price": inp.Price,
				"pricePlans.$[targetPlan].maxWeight": inp.MaxWeight,
				"pricePlans.$[targetPlan].minWeight": inp.MinWeight,
			},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: bson.A{bson.M{"targetPlan._id": inp.PlanID}},
		}),
	)
	if err != nil {
		return nil, DBError
	}
	return &Res{Messages.SuccessMsg, Messages.SuccessCode}, nil
}