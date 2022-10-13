package address

import "context"

func SubmitEditAddress(ctx context.Context, inp SubmitEditAddressInput) *AddressObj {
	initialObj := &AddressObj{
		UserPhone:         "09107655173",
		AddressInput:      inp.AddressInput,
		SelectedAddressID: inp.SelectedAddressID,
	}
	updateSelectedAddress := UpdateSelectedAddress{}
	validateAddingAddress := ValidateAddingAddress{NextChain: updateSelectedAddress}
	checkCustomerEssentialInfo := CheckCustomerEssentialInfo{NextChain: validateAddingAddress}
	genRawForm := GenRawForm{NextChain: checkCustomerEssentialInfo}
	res := genRawForm.Next(initialObj)
	return res
}
