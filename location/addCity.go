package location

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddCity(ctx context.Context, inp models.City)(*Res, error){
	var insertingCity models.City
	nameErr := _validateCityName(ctx, inp.Name, primitive.NilObjectID)
	if nameErr != nil {
		return nil, nameErr
	}
	provinceErr := _validateCityProvince(ctx, inp.ProvinceID)
	if provinceErr != nil {
		return nil, provinceErr
	}
	insertingCity.ID = primitive.NewObjectID()
	insertingCity.Name = inp.Name
	insertingCity.ProvinceID = inp.ProvinceID

	_, err := connection.City.InsertOne(ctx, insertingCity)
	if err != nil {
		return nil, InsertionError
	}
	return &Res{Messages.SuccessMsg, Messages.SuccessCode}, nil
}