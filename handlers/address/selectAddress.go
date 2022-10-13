package addressHndl

import (
	"github.com/amir5li/shipment/address"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SelectAddress(c *gin.Context) {
	var body address.SelectAddressInput
	c.Bind(&body)
	res := address.SelectAddress(body)
	session := sessions.Default(c)
	session.Set("addrID", res.SessionAddressID.Hex())
	session.Save()
	c.JSON(http.StatusOK, res.ConciseAddresses)
}
