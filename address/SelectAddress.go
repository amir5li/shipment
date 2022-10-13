package address

func SelectAddress(inp SelectAddressInput) *AddressObj {
	initialObj := &AddressObj{
		UserPhone:         "09107655173",
		SelectedAddressID: inp.AddressID,
	}
	setTargetForAddressesAndSession := SetTargetForAddressesAndSession{}
	setSelectedAddressAsDefault := SetSelectedAddressAsDefault{NextChain: setTargetForAddressesAndSession}
	getCustomerAddresses := GetCustomerAddresses{NextChain: setSelectedAddressAsDefault}
	res := getCustomerAddresses.Next(initialObj)
	return res
}
