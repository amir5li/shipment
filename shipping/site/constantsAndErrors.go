package shippingSite

type ShippingSiteErr string

func (se ShippingSiteErr) Error() string {
	return string(se)
}

const (
	InitialInfoTimeout ShippingSiteErr = "timeout in getting initial info"
)
