package address

type ValidateAddingAddress struct {
	NextChain AddressChain
}

func (va ValidateAddingAddress) Next(obj *AddressObj) *AddressObj {
	type fieldData struct{
		FieldName string
		ErrMsg *string
		Value interface{}

	}
	var fiedsInfo []fieldData
	// validate customer fields
	if obj.NeedUpdateCustomerInfo {
		// customer first name
		firstNameErr := _validateRawText(obj.AddressInput.CustomerFirstName)
		var field = fieldData{
			FieldName: CustomerFirstNameFieldName,
			Value: obj.AddressInput.CustomerFirstName,
		}
		if firstNameErr != nil {
			errMsg := firstNameErr.Error()
			field.ErrMsg = &errMsg
		}
		fiedsInfo = append(fiedsInfo, field)
		// customer last name
		lastNameErr := _validateRawText(obj.AddressInput.CustomerLastName)
		field = fieldData{
			FieldName: CustomerLastNameFieldName,
			Value: obj.AddressInput.CustomerLastName,
		}
		if lastNameErr != nil {
			errMsg := lastNameErr.Error()
			field.ErrMsg = &errMsg
		}
		fiedsInfo = append(fiedsInfo, field)
		// customer national code
		nationalCodeErr := _validateNationalCodeCustomer(obj.AddressInput.CustomerNationalCode)
		field = fieldData{
			FieldName: CustomerNationalCodeFieldName,
			Value: obj.AddressInput.CustomerNationalCode,
		}
		if nationalCodeErr != nil {
			errMsg := nationalCodeErr.Error()
			field.ErrMsg = &errMsg
		}
		fiedsInfo = append(fiedsInfo, field)
	}
	// consignee is customer
	if obj.AddressInput.ConsigneeIsCustomer {
		if obj.NeedUpdateCustomerInfo {
			for _, field := range fiedsInfo {
				switch field.FieldName{
				case CustomerFirstNameFieldName:
					var newField = field
					newField.FieldName = ConsigneeFirstNameFieldName
				}
			}
		}else{
			for _, sec := range obj.Form{
				if sec.Title == ConsigneeSectionTitle {
					for _, field := range sec.Fields {
						switch field.Name {
						case ConsigneeFirstNameFieldName:
							field.Value = obj.CustomerInfo.FirstName
						case ConsigneeLastNameFieldName:
							field.Value = obj.CustomerInfo.LastName
						case ConsigneeNationalCodeFieldName:
							field.Value = obj.CustomerInfo.NationalCode
						case ConsigneePhoneFieldName:
							field.Value = obj.CustomerInfo.Phone
						}
					}
				}
			}
		}
	}else{
		firstNameErr := _validateRawText(obj.AddressInput.ConsigneeFirstName)
		field := fieldData{
			FieldName: ConsigneeFirstNameFieldName,
			Value: obj.AddressInput.ConsigneeFirstName,
		}
		if firstNameErr != nil {
			errMsg := firstNameErr.Error()
			field.ErrMsg = &errMsg
		}
		fiedsInfo = append(fiedsInfo, field)
		lastNameErr := _validateRawText(obj.AddressInput.ConsigneeLastName)
		field = fieldData{
			FieldName: ConsigneeLastNameFieldName,
			Value: obj.AddressInput.ConsigneeLastName,
		}
		if lastNameErr != nil {
			errMsg := lastNameErr.Error()
			field.ErrMsg = &errMsg
		}
		fiedsInfo = append(fiedsInfo, field)
	}
}