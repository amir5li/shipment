package shippingSite

import (
	"context"
	"time"

	"github.com/amir5li/shipment/shipping/site/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getValidMethods(ctx context.Context, customerID primitive.ObjectID, addressID primitive.ObjectID) ([]ValidMethod, error) {
	var initialInfo struct {
		CityID      primitive.ObjectID
		TotalWeight uint
	}
	var weightChan = make(chan uint)
	var cityIDChan = make(chan primitive.ObjectID)
	func() {
		cityID := providers.CalculateCityOfSelectedAddress(ctx, customerID, addressID)
		cityIDChan <- cityID
	}()
	func() {
		totalWeight := providers.CalculateBasketWeight(ctx, customerID)
		weightChan <- totalWeight
	}()
	var receivedChans byte
	for {
		select {
		case cityID := <-cityIDChan:
			initialInfo.CityID = cityID
			receivedChans++
			if receivedChans == 2 {
				break
			}
		case totalWeight := <-weightChan:
			initialInfo.TotalWeight = totalWeight
			receivedChans++
			if receivedChans == 2 {
				break
			}
		case <-time.After(200 * time.Millisecond):
			return nil, InitialInfoTimeout
		}
	}
	validMethods := providers.GetValidMethods(ctx, initialInfo.CityID, initialInfo.TotalWeight)
	return validMethods, nil
}
