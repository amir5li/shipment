package shippingAdmin

import (
	"context"
	"fmt"
	"time"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func _validateNameOrTitle(text string, methodID primitive.ObjectID) error {
	if size := len(text); size < 4 {
		return InvalidLengthNameOrTitle
	}
	var ignoredID = methodID
	if ignoredID.IsZero() {
		ignoredID = primitive.NewObjectID()
	}
	count ,_ := connection.ShippingMethod.CountDocuments(
		context.TODO(),
		bson.M{
			"_id": bson.M{"$not": bson.M{"$eq": ignoredID}},
			"$or": bson.A{
				bson.M{"name": text},
				bson.M{"title": text},
			},
		},
	)
	if count != 0 {
		return DuplicateNameOrTitle
	}
	return nil
}

func _validateMinWeight(weight uint) error {
	if weight == 0 {
		return InvalidMinWeigth
	}
	return nil
}

func _validateMaxWeight(weight uint) error {
	if weight == 0 {
		return InvalidMaxWeigth
	}
	return nil
}

func _validatePrice(price uint) error {
	if price == 0 {
		return InvalidPrice
	}
	return nil
}

func _validateTimingBasket(hour uint, minute uint, pairedHour uint, pairedMinute uint,methodID primitive.ObjectID, weekday uint, basketID string) error {
	var aggRes []struct{
		ShipDayBaskets []models.ShippingBasket `bson:"shipDayBaskets"`
	}
	aggDecodings, _ := connection.ShippingMethod.Aggregate(
		context.TODO(),
		bson.A{
			bson.M{
				"$match": bson.M{
					"_id": methodID,
				},
			},
			bson.M{
				"$set": bson.M{
					"shipDayBaskets": bson.M{
						"$reduce": bson.M{
							"input": "$shippingDays",
							"initialValue": nil,
							"in": bson.M{
								"$cond": bson.M{
									"if": bson.M{
										"$and": bson.A{
											bson.M{
												"$eq": bson.A{"$$value", nil},
											},
											bson.M{
												"$eq": bson.A{"$$this.weekday", weekday},
											},
										},
									},
									"then": "$$this.baskets",
									"else": "$$value",
								},
							},
						},
					},
				},
			},
			bson.M{
				"$project": bson.M{
					"_id": 0,
					"shipDayBaskets": 1,
				},
			},
		},
	)
	aggDecodings.All(context.TODO(), &aggRes)
	baskets := aggRes[0].ShipDayBaskets
	filteredBasket := []models.ShippingBasket{}
	for _, basket := range baskets {
		if basket.ID != basketID{
			filteredBasket = append(filteredBasket, basket)
		}
	} 
	tarDur,_ := time.ParseDuration(fmt.Sprintf("%dh%dm", hour, minute))
	pairedDur, _ := time.ParseDuration(fmt.Sprintf("%dh%dm", pairedHour, pairedMinute))
	for _, basket := range filteredBasket {
		startDur, _ := time.ParseDuration(fmt.Sprintf("%dh%dm", basket.StartHour, basket.StartMinute))
		endDur, _ := time.ParseDuration(fmt.Sprintf("%dh%dm", basket.EndHour, basket.EndMinute))
		fmt.Println(tarDur.Minutes(), startDur.Minutes(), endDur.Minutes())
		if tarDur.Minutes() > startDur.Minutes() && tarDur.Minutes() < endDur.Minutes() {
			return ConflictTimingBasket
		}
		if (tarDur.Minutes() == startDur.Minutes() && pairedDur.Minutes() == endDur.Minutes()) || 
			(tarDur.Minutes() == endDur.Minutes() && pairedDur.Minutes() == startDur.Minutes()){
				return ConflictTimingBasket
			}
	}
	return nil
}

func _validatePriority(p uint) error{
	cnt, _ := connection.ShippingMethod.CountDocuments(context.TODO(), bson.M{"priority": p})
	if cnt != 0 {
		return DuplicatePriority
	}
	return nil
}