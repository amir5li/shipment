package location

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateCity(ctx context.Context, inp models.City)(*Res, error){
	nameErr := _validateCityName(ctx, inp.Name, inp.ID)
	if nameErr != nil {
		return nil, nameErr
	}
	_, err := connection.City.UpdateByID(
		ctx,
		inp.ID,
		bson.M{
			"$set": bson.M{
				"name": inp.Name,
			},
		},
	)
	if err != nil {
		return nil, UpdateError
	}
	return &Res{Messages.SuccessMsg, Messages.SuccessCode}, nil
}