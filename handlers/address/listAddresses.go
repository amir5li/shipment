package addressHndl

import (
	"github.com/amir5li/shipment/address"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func ListAddresses(c *gin.Context) {
	session := sessions.Default(c)
	addrHex := session.Get("addrID")
	var selectedAddress primitive.ObjectID
	if addrHex != nil {
		if primitive.IsValidObjectID(addrHex.(string)) {
			selectedAddress, _ = primitive.ObjectIDFromHex(addrHex.(string))
		}
	}
	res := address.ListAddresses(selectedAddress)
	if !res.SessionAddressID.IsZero() {
		session.Set("addrID", res.SessionAddressID.Hex())
		session.Save()
	}
	c.JSON(http.StatusOK, res.ConciseAddresses)
}
