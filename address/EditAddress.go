package address

import "context"

func EditAddress(ctx context.Context, inp EditAddressInput) *AddressObj {
	initialObj := &AddressObj{
		UserPhone: "09107655173",
		SelectedAddressID: inp.SelectedAddress,
	}
	fillFormForSelectedAddress := FillFormForSelectedAddress{}
	omitCustomerFormSection := OmitCustomerFormSection{NextChain: fillFormForSelectedAddress}
	genRawForm := GenRawForm{NextChain: omitCustomerFormSection}
	res := genRawForm.Next(initialObj)
	return res
}