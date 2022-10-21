package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PricePlan struct {
	ID        primitive.ObjectID `bson:"_id"`
	MaxWeight uint               `bson:"maxWeight"`
	MinWeight uint               `bson:"minWeight"`
	MinHeigth uint               `bson:"minHeigth"`
	MaxHeigth uint               `bson:"maxHeigth"`
	MinWidth  uint               `bson:"minWidth"`
	MaxWidth  uint               `bson:"maxWidth"`
	Price     uint               `bson:"price"`
}

type ShippingBasket struct {
	ID                string `bson:"id"`
	StartHour         uint   `bson:"startHour"`
	StartMinute       uint   `bson:"startMinute"`
	EndHour           uint   `bson:"endHour"`
	EndMinute         uint   `bson:"endMinute"`
	PreparationHour   uint   `bson:"preparationHour"`
	PreparationMinute uint   `bson:"preparationMinute"`
	Active            bool   `bson:"active"`
}

type ShippingDay struct {
	Name    string           `bson:"name"`
	Title   string           `bson:"title"`
	Weekday uint             `bson:"weekday"`
	Baskets []ShippingBasket `bson:"baskets"`
}
type ShipmentMethod struct {
	ID           primitive.ObjectID   `bson:"_id"`
	Name         string               `bson:"name"`
	Title        string               `bson:"title"`
	Priority     uint                 `bson:"priority"`
	Physical     bool                 `bson:"physical"`
	Description  string               `bson:"description"`
	ValidCities  []primitive.ObjectID `bson:"validCities"`
	PricePlans   []PricePlan          `bson:"pricePlans"`
	ShippingDays []ShippingDay        `bson:"shippingDays"`
}
