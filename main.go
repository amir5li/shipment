package main

import (
	addressHndl "github.com/amir5li/shipment/handlers/address"
	"github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"

	// locationHndl "github.com/amir5li/shipment/handlers/location"
	"github.com/gin-gonic/gin"
)

func main(){
	// route := gin.Default()
	// loc := route.Group("/loc")
	// {
	// 	loc.POST("/addCity", locationHndl.AddCity)
	// 	loc.POST("/addProvince", locationHndl.AddProvince)
	// 	loc.POST("/updateProvince", locationHndl.UpdateProvince)
	// 	loc.POST("/updateCity", locationHndl.UpdateCity)
	// 	loc.GET("/provinceList", locationHndl.GetProvinceList)
	// 	loc.POST("/cityList", locationHndl.GetCityList)
	// }
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379","", []byte("secret"))
	// store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("customerInfo", store))
	addr := r.Group("/address")
	{
		addr.GET("/add", addressHndl.AddAddress)
		addr.POST("/submit-add", addressHndl.SubmitAddAddress)
		addr.POST("/edit", addressHndl.EditAddress)
	}
	// route.Run(":6000")
	r.Run(":6500")
}