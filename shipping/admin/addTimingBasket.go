package shippingAdmin

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddTimingBasket(ctx context.Context, inp AddTimingBasketInput)(*Res, error){
	startDurErr := _validateTimingBasket(inp.StartHour, inp.StartMinute, inp.EndHour, inp.EndMinute, inp.MethodID, inp.Weekday, string(""))
	if startDurErr != nil {
		return nil, startDurErr
	}
	endDurError := _validateTimingBasket(inp.EndHour, inp.EndMinute, inp.StartHour, inp.StartMinute, inp.MethodID, inp.Weekday, string(""))
	if endDurError != nil {
		return nil, endDurError
	}
	insertingTimingBasket := models.ShippingBasket{
		ID: uuid.NewString(),
		StartHour: inp.StartHour,
		StartMinute: inp.StartMinute,
		EndHour: inp.EndHour,
		EndMinute: inp.EndMinute,
		PreparationHour: inp.PreparationHour,
		PreparationMinute: inp.PreparationMinute,
		Active: inp.Active,
	}
	_, err := connection.ShippingMethod.UpdateByID(
		ctx,
		inp.MethodID,
		bson.M{
			"$push": bson.M{
				"shippingDays.$[targetDay].baskets": bson.M{
					"$each": bson.A{insertingTimingBasket},
					"$sort": bson.M{"startHour": 1},
				},
			},
		},
		options.Update().SetArrayFilters(options.ArrayFilters{
			Filters: bson.A{bson.M{"targetDay.weekday": inp.Weekday}},
		}),
	)
	if err != nil {
		return nil, DBError
	}
	return &Res{Messages.SuccessMsg, Messages.SuccessCode}, nil
}