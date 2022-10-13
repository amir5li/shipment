package addressHndl

import (
	"net/http"

	"github.com/amir5li/shipment/address"
	"github.com/gin-gonic/gin"
)

func EditAddress(c *gin.Context){
	var body address.EditAddressInput
	c.Bind(&body)
	res := address.EditAddress(c, body)
	c.JSON(http.StatusOK, res)
}