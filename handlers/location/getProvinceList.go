package locationHndl

import (
	"net/http"

	"github.com/amir5li/shipment/location"
	"github.com/gin-gonic/gin"
)

func GetProvinceList(c *gin.Context){
	res := location.GetProvinceList(c)
	c.JSON(http.StatusOK, res)
}