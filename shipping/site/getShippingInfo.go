package shippingSite

import (
	"fmt"

	"github.com/amir5li/shipment/shipping/site/providers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetShippingInfo(c *gin.Context, inp GetShippingDataInput)(interface{}, error){
	session := sessions.Default(c)
	addrIDHex := session.Get("addrID")
	var addrID primitive.ObjectID
	if addrIDHex == nil {
		return nil, AddressNotSelected
	}
	addrID, err := primitive.ObjectIDFromHex(addrIDHex.(string))
	testCustomerID, _ := primitive.ObjectIDFromHex("62952f08b718fa718218f170")
	fmt.Println(addrID, err)
	methods, err := providers.GetValidMethods(c, testCustomerID, addrID)
	if err != nil {
		return nil, err
	}
	fmt.Println("methods: ", methods)
	return nil, nil
}