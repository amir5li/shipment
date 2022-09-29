package location

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func _validateCityName(ctx context.Context, name string, cityID primitive.ObjectID)error{
	var ignoreID primitive.ObjectID = cityID
	if cityID.Hex() == primitive.NilObjectID.Hex() {
		ignoreID = primitive.NewObjectID()
	}
	count, _ := connection.City.CountDocuments(ctx, bson.M{"_id": bson.M{"$ne": ignoreID}, "name": name})
	if count != 0 {
		return DuplicateLocationName
	}
	return nil
}

func _validateCityProvince(ctx context.Context, provinceID primitive.ObjectID) error {
	if !primitive.IsValidObjectID(provinceID.Hex()) || provinceID.IsZero() {
		return InvalidProvinceID
	}
	return nil
}

func _validateProvinceName(ctx context.Context, name string, provinceID primitive.ObjectID) error {
	var ignoreID = provinceID
	if provinceID.Hex() == primitive.NilObjectID.Hex() {
		ignoreID = primitive.NewObjectID()
	}
	count, _ := connection.Province.CountDocuments(ctx, bson.M{"_id": bson.M{"$ne": ignoreID}, "name": name})
	if count != 0 {
		return DuplicateLocationName
	}
	return nil
}
