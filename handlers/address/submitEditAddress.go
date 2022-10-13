package addressHndl

import (
	"github.com/amir5li/shipment/address"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SubmitEditAddress(c *gin.Context) {
	var body address.SubmitEditAddressInput
	c.Bind(&body)
	res := address.SubmitEditAddress(c, body)
	c.JSON(http.StatusOK, res)
}
