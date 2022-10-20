package shippingSiteHndl

import (
	"net/http"

	"github.com/amir5li/shipment/shipping/site"
	"github.com/gin-gonic/gin"
)

func GetShippingInfo(c *gin.Context){
	var body shippingSite.GetShippingDataInput
	c.Bind(&body)
	res, err := shippingSite.GetShippingInfo(c, body)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error(), "code": -1})
		return
	}
	c.JSON(http.StatusOK, res)
}