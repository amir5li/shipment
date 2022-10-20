package main

import (
	"log"
	"net/http"
	"time"

	addressHndl "github.com/amir5li/shipment/handlers/address"
	"github.com/amir5li/shipment/handlers/shipping/admin"
	"github.com/amir5li/shipment/handlers/shipping/site"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)
var g errgroup.Group
func siteRouter() http.Handler{
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	// store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("customerInfo", store))
	addr := r.Group("/address")
	{
		addr.GET("/add", addressHndl.AddAddress)
		addr.POST("/submit-add", addressHndl.SubmitAddAddress)
		addr.POST("/edit", addressHndl.EditAddress)
		addr.POST("/submit-edit", addressHndl.SubmitEditAddress)
		addr.POST("/select", addressHndl.SelectAddress)
		addr.GET("/list", addressHndl.ListAddresses)
	}
	shipping := r.Group("/shipping")
	{
		shipping.POST("/get", shippingSiteHndl.GetShippingInfo)
	}
	return r
}
func adminRouter() http.Handler {
	r := gin.Default()
	shipment := r.Group("/shipment")
	{
		shipment.POST("/create", shippingAdminHndl.CreateShippingMethod)
		shipment.POST("/add-plan", shippingAdminHndl.AddPricePlan)
		shipment.POST("/edit-plan", shippingAdminHndl.EditPricePlan)
		shipment.POST("/delete-plan", shippingAdminHndl.DeletePricePlan)
		shipment.POST("/add-basket", shippingAdminHndl.AddTimingBasket)
		shipment.POST("/edit-basket", shippingAdminHndl.EditTimingBasket)
		shipment.POST("/delete-basket", shippingAdminHndl.DeleteTimingBasket)
	}
	return r

}
func main() {
	siteServer := &http.Server{
		Addr: ":6500",
		Handler: siteRouter(),
		WriteTimeout: 5 * time.Second,
		ReadTimeout: 5 * time.Second,
	}
	adminServer := &http.Server{
		Addr: ":6800",
		Handler: adminRouter(),
		WriteTimeout: 5 * time.Second,
		ReadTimeout: 5 * time.Second,
	}
	g.Go(func() error {
       err := siteServer.ListenAndServe()
	   if err != nil && err != http.ErrServerClosed {
		   log.Fatal(err)
	   }
	   return err
	})
	g.Go(func() error {
		err := adminServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
