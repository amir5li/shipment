package shippingAdmin

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"github.com/amir5li/shipment/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateShippingMethod(ctx context.Context, inp CreateMethodInput) (*Res, error) {
	nameErr := _validateNameOrTitle(inp.Name, primitive.NilObjectID)
	if nameErr != nil {
		return nil, nameErr
	}
	titleErr := _validateNameOrTitle(inp.Title, primitive.NilObjectID)
	if titleErr != nil {
		return nil, titleErr
	}
	priorityErr := _validatePriority(inp.Priority)
	if priorityErr != nil {
		return nil, priorityErr
	}
	insertingMethod := models.ShipmentMethod{
		ID:           primitive.NewObjectID(),
		Name:         inp.Name,
		Title:        inp.Title,
		Priority:     inp.Priority,
		Description:  inp.Description,
		Physical:     inp.Physical,
		ValidCities:  []primitive.ObjectID{},
		PricePlans:   []models.PricePlan{},
		ShippingDays: initialShippingDays,
	}
	_, err := connection.ShippingMethod.InsertOne(
		ctx,
		insertingMethod,
	)
	if err != nil {
		return nil, DBError
	}
	return &Res{Messages.SuccessMsg, Messages.SuccessCode}, nil
}
