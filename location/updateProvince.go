package location

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateProvince(ctx context.Context, inp models.Province)(*Res, error){
	nameErr := _validateProvinceName(ctx, inp.Name, inp.ID)
	if nameErr != nil {
		return nil, nameErr
	}
	_, err := connection.Province.UpdateByID(
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