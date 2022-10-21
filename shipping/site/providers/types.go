package providers

import (
	"github.com/amir5li/shipment/models"
)

type ValidMethod struct {
	Title    string
	Priority uint
	Price    uint
	Physical bool
	Days     []models.ShippingDay
}

type SelectedMethods struct {
	Regular  *ValidMethod
	Physical *ValidMethod
	Overall  *ValidMethod
}
