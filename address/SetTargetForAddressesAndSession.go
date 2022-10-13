package address

type SetTargetForAddressesAndSession struct {
	NextChain AddressChain
}

func (stas SetTargetForAddressesAndSession) Next(obj *AddressObj) *AddressObj {
	for _, addr := range obj.ConciseAddresses {
		if addr.ID.Hex() == obj.SelectedAddressID.Hex() {
			addr.Selected = true
		}
	}
	obj.SessionAddressID = obj.SelectedAddressID
	if stas.NextChain != nil {
		newObj := stas.NextChain.Next(obj)
		return newObj
	}
	return obj
}
