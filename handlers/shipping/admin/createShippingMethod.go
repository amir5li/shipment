package shippingAdminHndl

import (
	"net/http"

	"github.com/amir5li/shipment/shipping/admin"
	"github.com/gin-gonic/gin"
)

func CreateShippingMethod(c *gin.Context){
	var body shippingAdmin.CreateMethodInput
	c.Bind(&body)
	res, err := shippingAdmin.CreateShippingMethod(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error(), "code": -1})
		return
	}
	c.JSON(http.StatusOK, res)
}