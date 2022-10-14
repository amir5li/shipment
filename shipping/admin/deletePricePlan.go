package shippingAdmin

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
)

func DeletePricePlan(ctx context.Context, inp DeletePricePlanInput) (*Res, error){
	_, err := connection.ShippingMethod.UpdateByID(
		ctx,
		inp.MethodID,
		bson.M{
			"$pull": bson.M{
				"pricePlans": bson.M{
					"_id": inp.PlanID,
				},
			},
		},
	)
	if err != nil {
		return nil, DBError
	}
	return &Res{Messages.SuccessMsg, Messages.SuccessCode}, nil
}