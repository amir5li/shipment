package address



type AddressChain interface {
	Next(*AddressObj) *AddressObj 
}