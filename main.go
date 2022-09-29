package main

import (
	locationHndl "github.com/amir5li/shipment/handlers/location"
	"github.com/gin-gonic/gin"
)

func main(){
	route := gin.Default()
	loc := route.Group("/loc")
	{
		loc.POST("/addCity", locationHndl.AddCity)
		loc.POST("/addProvince", locationHndl.AddProvince)
		loc.POST("/updateProvince", locationHndl.UpdateProvince)
		loc.POST("/updateCity", locationHndl.UpdateCity)
		loc.GET("/provinceList", locationHndl.GetProvinceList)
		loc.POST("/cityList", locationHndl.GetCityList)
	}
	route.Run(":6000")
}