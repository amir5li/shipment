package providers

import (
	"github.com/amir5li/shipment/models"
)

type ValidMethod struct {
	Title string
	Priority uint
	Price uint
	Days []models.ShippingDay
}

