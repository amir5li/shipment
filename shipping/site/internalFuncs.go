package shippingSite

import (
	"context"

	"github.com/amir5li/shipment/shipping/site/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getValidMethods(ctx context.Context, customerID primitive.ObjectID, addressID primitive.ObjectID){
	var initialInfo struct {
		CityID primitive.ObjectID
		TotalWeight uint
	}
	var weightChan = make(chan uint)
	var cityIDChan = make(chan primitive.ObjectID)
	func(){
		cityID := providers.CalculateCityOfSelectedAddress(ctx, customerID, addressID)
		cityIDChan <- cityID
	}()
	func(){
		totalWeight := providers.CalculateBasketWeight(ctx)
	}()
}