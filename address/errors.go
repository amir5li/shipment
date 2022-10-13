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
	InvalidPostalAddress AddressErr = "postal address can just include persian alphabetic digits comma and paranteses"
	InvalidPostalCode AddressErr = "invalid postal code"
	InvalidCity AddressErr = "city must be selected"
	InvalidProvince AddressErr = "province must be selected"
)