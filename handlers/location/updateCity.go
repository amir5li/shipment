package locationHndl

import (
	"net/http"

	"github.com/amir5li/shipment/location"
	"github.com/amir5li/shipment/models"
	"github.com/gin-gonic/gin"
)

func UpdateCity(c *gin.Context){
	var body models.City
	c.Bind(&body)
	res, err := location.UpdateCity(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error(), "code": -1})
		return
	}
	c.JSON(http.StatusOK, res)
}