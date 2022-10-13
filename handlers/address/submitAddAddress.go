package addressHndl

import (
	"fmt"
	"net/http"

	"github.com/amir5li/shipment/address"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SubmitAddAddress(c *gin.Context) {
	session := sessions.Default(c)
	addrID := session.Get("addrID")
	fmt.Println("myAddrID", addrID)
	var body address.AddressInput
	c.Bind(&body)
	res := address.SubmitAddAddress(c, body)
	fmt.Println(res.SessionAddressID)
	if res.SessionAddressID.Hex() != primitive.NilObjectID.Hex() {
		fmt.Println("im called")
		session.Set("addrID", res.SessionAddressID.Hex())
		session.Save()
	}
	c.JSON(http.StatusOK, res)
}
