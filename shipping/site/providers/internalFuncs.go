package providers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetValidMethods(ctx context.Context, customerID primitive.ObjectID, addressID primitive.ObjectID) ([]ValidMethod, error) {
	var initialInfo struct {
		CityID      primitive.ObjectID
		TotalWeight uint
	}
	var weightChan = make(chan uint)
	var cityIDChan = make(chan primitive.ObjectID)
	go func() {
		fmt.Print("im call")
		cityID := calculateCityOfSelectedAddress(ctx, customerID, addressID)
		cityIDChan <- cityID
	}()
	go func() {
		totalWeight := calculateBasketWeight(ctx, customerID)
		weightChan <- totalWeight
	}()
	var receivedChans byte
	chans:
	for {
		select {
		case cityID := <-cityIDChan:
			initialInfo.CityID = cityID
			receivedChans++
			if receivedChans == 2 {
				break chans
			}
		case totalWeight := <-weightChan:
			initialInfo.TotalWeight = totalWeight
			receivedChans++
			if receivedChans == 2 {
				break chans
			}
		case <-time.After(200 * time.Millisecond):
			return nil, InitialInfoTimeout
		}
	}
	validMethods := getValidMethods(ctx, initialInfo.CityID, initialInfo.TotalWeight)
	return validMethods, nil
}
