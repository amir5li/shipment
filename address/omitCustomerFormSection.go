package address

type OmitCustomerFormSection struct {
	NextChain AddressChain
}

func (ocs OmitCustomerFormSection) Next(obj *AddressObj) *AddressObj {
	var newForm []*AddressSection
	for _, sec := range obj.Form {
		if sec.Title != CustomerSectionTitle {
			newForm = append(newForm, sec)
		}
	}
	obj.Form = newForm
	if ocs.NextChain != nil {
		newObj := ocs.NextChain.Next(obj)
		return newObj
	}
	return obj
}
