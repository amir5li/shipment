package location

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProvince(ctx context.Context, inp models.Province)(*Res, error){
	nameErr := _validateProvinceName(ctx, inp.Name, primitive.NilObjectID)
	if nameErr != nil {
		return nil, nameErr
	}
	var insertingProvince models.Province
	insertingProvince.ID = primitive.NewObjectID()
	insertingProvince.Name = inp.Name
	_, err := connection.Province.InsertOne(ctx, insertingProvince)
	if err != nil {
		return nil, InsertionError
	}
	return &Res{Messages.SuccessMsg, Messages.SuccessCode}, nil
}