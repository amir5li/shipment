package locationHndl

import (
	"net/http"
	"github.com/amir5li/shipment/location"
	"github.com/amir5li/shipment/models"
	"github.com/gin-gonic/gin"
)

func UpdateProvince(c *gin.Context){
	var body models.Province
	c.Bind(&body)
	res, err := location.UpdateProvince(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error(), "code": -1})
		return
	}
	c.JSON(http.StatusOK, res)
}