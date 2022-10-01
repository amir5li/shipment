package address

import (
	"context"
)

func AddAddress(ctx context.Context) *AddressObj {
	testPhone := "09107655173"
	obj := &AddressObj{
		UserPhone: testPhone,
	}
	checkCustomerInfo := CheckCustomerEssentialInfo{}
	genForm := GenRawForm{NextChain: checkCustomerInfo}
	res := genForm.Next(obj)
	return res
}
