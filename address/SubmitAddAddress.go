package address

import "context"

func SubmitAddAddress(ctx context.Context, addrInput AddressInput)*AddressObj{
	initialObj := &AddressObj{
		UserPhone: "09107655173",
		AddressInput: addrInput,
	}
	getCustomerAddresses := GetCustomerAddresses{}
	setNewAddressDefault := SetNewAddressDefaltIsNeeded{NextChain: getCustomerAddresses}
	addNewAddress := AddNewAddress{NextChain: setNewAddressDefault}
	updateCustomerInfo := UpdateCustomerInfo{NextChan: addNewAddress}
	validateNewAddr := ValidateAddingAddress{NextChain: updateCustomerInfo}
	checkCustomerInfo := CheckCustomerEssentialInfo{NextChain: validateNewAddr}
	genForm := GenRawForm{NextChain: checkCustomerInfo}
	res :=genForm.Next(initialObj)
	return res
}