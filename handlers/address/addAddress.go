package addressHndl

import (
	"net/http"

	"github.com/amir5li/shipment/address"
	"github.com/gin-gonic/gin"
)

func AddAddress(c *gin.Context){
	res := address.AddAddress(c)
	c.JSON(http.StatusOK, res)
}