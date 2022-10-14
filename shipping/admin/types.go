package shippingAdmin

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateMethodInput struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type AddPricePlanInput struct {
	MethodID primitive.ObjectID `json:"methodID"`
	MaxWeight uint `json:"maxWeight"`
	MinWeight uint `json:"minWeight"`
	Price uint `json:"price"`
}

type EditPricePlanInput struct {
	MethodID primitive.ObjectID `json:"methodID"`
	PlanID primitive.ObjectID `json:"planID"`
	MaxWeight uint `json:"maxWeight"`
	MinWeight uint `json:"minWeight"`
	Price uint `json:"price"`
}

type DeletePricePlanInput struct {
	MethodID primitive.ObjectID `json:"methodID"`
	PlanID primitive.ObjectID `json:"planID"`
}

type AddTimingBasketInput struct {
	MethodID primitive.ObjectID `json:"methodID"`
	Weekday uint `json:"weekday"`
	StartHour uint `json:"startHour"`
	StartMinute uint `json:"startMinute"`
	EndHour uint `json:"endHour"`
	EndMinute uint `json:"endMinute"`
	PreparationHour uint `json:"preparationHour"`
	PreparationMinute uint `json:"preparationMinute"`
	Active bool `json:"active"`
}

type EditTimingBasketInput struct {
	MethodID primitive.ObjectID `json:"methodID"`
	BasketID string `json:"basketID"`
	Weekday uint `json:"weekday"`
	StartHour uint `json:"startHour"`
	StartMinute uint `json:"startMinute"`
	EndHour uint `json:"endHour"`
	EndMinute uint `json:"endMinute"`
	PreparationHour uint `json:"preparationHour"`
	PreparationMinute uint `json:"preparationMinute"`
	Active bool `json:"active"`
}

type DeleteTimingBasketInput struct {
	MethodID primitive.ObjectID `json:"methodID"`
	BasketID string `json:"basketID"`
	Weekday uint `json:"weekday"`
}

type UpdateValidCitiesInput struct {
	MethodID primitive.ObjectID `json:"methodID"`
	Cities []primitive.ObjectID `js`
}

type Res struct {
	Msg string `json:"msg"`
	Code byte `json:"code"`
}