package address

type AddressErr string

func (ae AddressErr) Error() string {
	return string(ae)
}
const (
	InvalidPhoneNumber AddressErr = "entered phone is invalid"
	InvalidText AddressErr = "entered name is invalid"
	InvalidNationalCode AddressErr = "entered national code is invalid"
	DuplicateNationalCode AddressErr = "entered national code is duplicate"
)