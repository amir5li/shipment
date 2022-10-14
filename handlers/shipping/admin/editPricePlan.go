package shippingAdminHndl

import (
	"net/http"

	shippingAdmin "github.com/amir5li/shipment/shipping/admin"
	"github.com/gin-gonic/gin"
)

func EditPricePlan(c *gin.Context){
	var body shippingAdmin.EditPricePlanInput
	c.Bind(&body)
	res, err := shippingAdmin.EditPricePlan(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error(), "code": -1})
		return
	}
	c.JSON(http.StatusOK, res)
}