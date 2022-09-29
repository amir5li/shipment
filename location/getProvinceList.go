package location

import (
	"context"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProvinceList(ctx context.Context) []Loc {
	var res []Loc
	aggDecoding, _ := connection.Province.Aggregate(
		ctx, 
		bson.A{},
	)
	aggDecoding.All(ctx, &res)
	return res
}