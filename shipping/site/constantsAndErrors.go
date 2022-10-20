package shippingSite
type ShippingSiteErr string

func (se ShippingSiteErr) Error() string {
	return string(se)
}

const (
	AddressNotSelected ShippingSiteErr = "first select your address"
)