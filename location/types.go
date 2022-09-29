package location

import "go.mongodb.org/mongo-driver/bson/primitive"

type Res struct {
	Msg string `json:"msg"`
	Code byte `json:"code"`
}
type GetCityListInput struct {
	ProvinceID primitive.ObjectID `json:"provinceID"`
}

type Loc struct {
	Name string `bson:"name" json:"name"`
	ID primitive.ObjectID `bson:"_id" json:"id"`
}