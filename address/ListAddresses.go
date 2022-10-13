package address

import "go.mongodb.org/mongo-driver/bson/primitive"

func ListAddresses(selectedAddr primitive.ObjectID) *AddressObj {
	initialObj := &AddressObj{
		UserPhone:         "09107655173",
		SelectedAddressID: selectedAddr,
	}
	findSelectedAddress := FindSelectedAddress{}
	getCustomerAddresses := GetCustomerAddresses{NextChain: findSelectedAddress}
	res := getCustomerAddresses.Next(initialObj)
	return res
}
