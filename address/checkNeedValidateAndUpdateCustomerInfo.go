package address

type CheckNeedUpdateCustomerInfo struct{
	NextChain AddressChain
}

func (nu CheckNeedUpdateCustomerInfo) Next(obj *AddressObj) *AddressObj {
	var needUpdateCustomerInfo bool
	for _, sec := range obj.Form {
		if sec.Title == CustomerSectionTitle {
			needUpdateCustomerInfo = true
		}
	}
	obj.NeedUpdateCustomerInfo = needUpdateCustomerInfo
	if nu.NextChain != nil {
		newObj := nu.NextChain.Next(obj)
		return newObj
	}
	return obj
}