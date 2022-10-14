package shippingAdmin

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DeleteTimingBasket(ctx context.Context, inp DeleteTimingBasketInput)(*Res, error){
	_, err := connection.ShippingMethod.UpdateByID(
		ctx,
		inp.MethodID,
		bson.M{
			"$pull": bson.M{
				"shippingDays.$[targetDay].baskets": bson.M{
					"id": inp.BasketID,
				},
			},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: bson.A{bson.M{"targetDay.weekday":inp.Weekday}},
		}),
	)
	if err != nil {
		return nil, DBError
	}
	return &Res{Messages.SuccessMsg, Messages.SuccessCode}, nil
}