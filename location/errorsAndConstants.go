package location

type LocationErr string

func (le LocationErr) Error() string {
	return string(le)
}

const (
	DuplicateLocationName LocationErr = "name exist on another instance"
	InvalidProvinceID LocationErr = "inserted provinceID is invalid"
	InsertionError LocationErr = "insertion error"
	UpdateError LocationErr = "update error"
)

var Messages = struct{
	SuccessMsg string
	SuccessCode byte
}{
	SuccessMsg: "operation done successfully",
	SuccessCode: 1,
}