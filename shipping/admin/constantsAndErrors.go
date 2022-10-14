package shippingAdmin

import "github.com/amir5li/shipment/models"

type ShippingAdminErr string

func (se ShippingAdminErr) Error() string {
	return string(se)
}

const (
	InvalidLengthNameOrTitle ShippingAdminErr = "the length of name or title must be atleast 4"
	DuplicateNameOrTitle ShippingAdminErr = "this name or title used in another shipping method"
	DBError ShippingAdminErr = "something unexpectional happened please try again"
	InvalidMinWeigth ShippingAdminErr = "invalid minum weight inserted"
	InvalidMaxWeigth ShippingAdminErr = "invalid maximum weight inserted"
	InvalidPrice ShippingAdminErr = "invalid price inserted"
	ConflictTimingBasket ShippingAdminErr = "inserted timing basket is invalid"
)

var Messages = struct {
	SuccessMsg string
	SuccessCode byte
}{
	SuccessMsg: "done",
	SuccessCode: 1,
}

var initialShippingDays = []models.ShippingDay{
	{
		Name: "Monday",
		Title: "دوشنبه",
		Weekday: 1,
		Baskets: []models.ShippingBasket{},
	},
	{
		Name: "Tuesday",
		Title: "سه شنبه",
		Weekday: 2,
		Baskets: []models.ShippingBasket{},
	},
	{
		Name: "Wednesday",
		Title: "چهارشنبه",
		Weekday: 3,
		Baskets: []models.ShippingBasket{},
	},
	{
		Name: "Thursday",
		Title: "پنجشنبه",
		Weekday: 4,
		Baskets: []models.ShippingBasket{},
	},
	{
		Name: "Friday",
		Title: "جمعه",
		Weekday: 5,
		Baskets: []models.ShippingBasket{},
	},
	{
		Name: "Saturday",
		Title: "شنبه",
		Weekday: 6,
		Baskets: []models.ShippingBasket{},
	},
	{
		Name: "Sunday",
		Title: "یکشنبه",
		Weekday: 7,
		Baskets: []models.ShippingBasket{},
	},
}