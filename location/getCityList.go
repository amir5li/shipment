package location

import (
	"context"
	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCityList(ctx context.Context, inp GetCityListInput) []Loc {
	matchBson := bson.M{}
	if !inp.ProvinceID.IsZero() {
		matchBson["provinceID"] = inp.ProvinceID
	}
	var res []Loc
	aggDecoding, _ := connection.City.Aggregate(
		ctx,
		bson.A{
			bson.M{"$match": matchBson},
		},
	)
	aggDecoding.All(ctx, &res)
	if res == nil {
		res = []Loc{}
	}
	return res
}