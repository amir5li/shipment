package shippingAdmin

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EditTimingBasket(ctx context.Context, inp EditTimingBasketInput)(*Res, error){
	startDurError := _validateTimingBasket(inp.StartHour, inp.StartMinute, inp.EndHour, inp.EndMinute, inp.MethodID, inp.Weekday, inp.BasketID)
	if startDurError != nil {
		return nil, startDurError
	}
	endDurError := _validateTimingBasket(inp.EndHour, inp.EndMinute, inp.StartHour, inp.StartMinute, inp.MethodID, inp.Weekday, inp.BasketID)
	if endDurError != nil {
		return nil, endDurError
	}
	_, err := connection.ShippingMethod.UpdateByID(
		ctx,
		inp.MethodID,
		bson.M{
			"$set": bson.M{
				"shippingDays.$[targetDay].baskets.$[targetBasket].startHour": inp.StartHour,
				"shippingDays.$[targetDay].baskets.$[targetBasket].startMinute": inp.StartMinute,
				"shippingDays.$[targetDay].baskets.$[targetBasket].endHour": inp.EndHour,
				"shippingDays.$[targetDay].baskets.$[targetBasket].endMinute": inp.EndMinute ,
				"shippingDays.$[targetDay].baskets.$[targetBasket].preparationHour": inp.PreparationHour,
				"shippingDays.$[targetDay].baskets.$[targetBasket].preparationMinute": inp.PreparationMinute,
				"shippingDays.$[targetDay].baskets.$[targetBasket].active": inp.Active,
			},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: bson.A{bson.M{"targetDay.weekday": inp.Weekday}, bson.M{"targetBasket.id": inp.BasketID}},
		}),
	)
	if err != nil {
		return nil, DBError
	}
	return &Res{Messages.SuccessMsg, Messages.SuccessCode}, nil
}