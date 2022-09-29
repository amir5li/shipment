package locationHndl

import (
	"net/http"

	"github.com/amir5li/shipment/location"
	"github.com/gin-gonic/gin"
)

func GetCityList(c *gin.Context){
	var body location.GetCityListInput
	c.Bind(&body)
	res := location.GetCityList(c, body)
	c.JSON(http.StatusOK, res)
}