package providers

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getValidMethods(ctx context.Context, cityID primitive.ObjectID, totalWeight uint)[]ValidMethod{
	var allMethods []models.ShipmentMethod
	aggDec, _ := connection.ShippingMethod.Aggregate(ctx, bson.A{})
	aggDec.All(ctx, &allMethods)
	var validMethods []ValidMethod
	for _, method := range allMethods {
		// validate city
		var canFindCity bool
		for _, city := range method.ValidCities {
			if city.Hex() == cityID.Hex(){
				canFindCity = true
			}
		}
		if !canFindCity {
			continue
		}
		// validate weight
		var canFindTargetWeight bool
		var matchedPrice uint
		for _, wp := range method.PricePlans {
			if totalWeight >= wp.MinWeight && totalWeight <= wp.MaxWeight {
				canFindTargetWeight = true
				matchedPrice = wp.Price
			}
		}
		if !canFindTargetWeight {
			continue
		}
		validMethods = append(validMethods, ValidMethod{
			Title: method.Title,
			Priority: method.Priority,
			Price: matchedPrice,
			Days: method.ShippingDays,
		})
	}
	return validMethods
}