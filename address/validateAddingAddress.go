package address

import "fmt"

type ValidateAddingAddress struct {
	NextChain AddressChain
}

func (va ValidateAddingAddress) Next(obj *AddressObj) *AddressObj {
	type fieldData struct {
		FieldName string
		ErrMsg    *string
		Value     interface{}
	}
	var fiedsInfo []fieldData
	// validate customer fields
	fmt.Println(obj.NeedUpdateCustomerInfo)
	if obj.NeedUpdateCustomerInfo {
		// customer first name
		firstNameErr := _validateRawText(obj.AddressInput.CustomerFirstName)
		fmt.Println("firstNameErr: ", firstNameErr)
		var field = fieldData{
			FieldName: CustomerFirstNameFieldName,
			Value:     obj.AddressInput.CustomerFirstName,
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
			Value:     obj.AddressInput.CustomerLastName,
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
			Value:     obj.AddressInput.CustomerNationalCode,
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
				switch field.FieldName {
				case CustomerFirstNameFieldName:
					var newField = field
					newField.FieldName = ConsigneeFirstNameFieldName
				}
			}
		} else {
			for _, sec := range obj.Form {
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
	} else {
		firstNameErr := _validateRawText(obj.AddressInput.ConsigneeFirstName)
		field := fieldData{
			FieldName: ConsigneeFirstNameFieldName,
			Value:     obj.AddressInput.ConsigneeFirstName,
		}
		if firstNameErr != nil {
			errMsg := firstNameErr.Error()
			field.ErrMsg = &errMsg
		}
		fiedsInfo = append(fiedsInfo, field)
		lastNameErr := _validateRawText(obj.AddressInput.ConsigneeLastName)
		field = fieldData{
			FieldName: ConsigneeLastNameFieldName,
			Value:     obj.AddressInput.ConsigneeLastName,
		}
		if lastNameErr != nil {
			errMsg := lastNameErr.Error()
			field.ErrMsg = &errMsg
		}
		fiedsInfo = append(fiedsInfo, field)
		nationalCodeErr := _validateNationalCodeConsignee(obj.AddressInput.ConsigneeNationalCode)
		field = fieldData{
			FieldName: ConsigneeNationalCodeFieldName,
			Value:     obj.AddressInput.ConsigneeNationalCode,
		}
		if nationalCodeErr != nil {
			errMsg := nationalCodeErr.Error()
			field.ErrMsg = &errMsg
		}
		fiedsInfo = append(fiedsInfo, field)
		phone, phoneErr := _validatePhone(obj.AddressInput.ConsigneePhone)
		field = fieldData{
			FieldName: ConsigneePhoneFieldName,
			Value:     phone,
		}
		if phoneErr != nil {
			errMsg := phoneErr.Error()
			field.ErrMsg = &errMsg
			field.Value = obj.AddressInput.ConsigneePhone
		}
		fiedsInfo = append(fiedsInfo, field)
	}
	// validate Address Input
	// address
	postalAddressErr := _validateAddressPostal(obj.AddressInput.AddressPostalAddress)
	field := fieldData{
		FieldName: AddressPostalAddressFieldName,
		Value:     obj.AddressInput.AddressPostalAddress,
	}
	if postalAddressErr != nil {
		errMsg := postalAddressErr.Error()
		field.ErrMsg = &errMsg
	}
	fiedsInfo = append(fiedsInfo, field)
	// postal code
	postalCodeErr := _validateAddressPostalCode(obj.AddressInput.AddressPostalCode)
	field = fieldData{
		FieldName: AddressPostalCodeFieldName,
		Value:     obj.AddressInput.AddressPostalCode,
	}
	if postalAddressErr != nil {
		errMsg := postalCodeErr.Error()
		field.ErrMsg = &errMsg
	}
	fiedsInfo = append(fiedsInfo, field)
	// province
	provinceErr := _validateAddressProvince(obj.AddressInput.AddressProvince)
	field = fieldData{
		FieldName: AddressProvinceFieldName,
		Value:     obj.AddressInput.AddressProvince,
	}
	if provinceErr != nil {
		errMsg := provinceErr.Error()
		field.ErrMsg = &errMsg
	}
	fiedsInfo = append(fiedsInfo, field)
	// city
	cityErr := _validateAddressCity(obj.AddressInput.AddressCity, obj.AddressInput.AddressProvince)
	field = fieldData{
		FieldName: AddressCityFieldName,
		Value:     obj.AddressInput.AddressCity,
	}
	if cityErr != nil {
		errMsg := cityErr.Error()
		field.ErrMsg = &errMsg
	}
	fiedsInfo = append(fiedsInfo, field)
	// unit
	field = fieldData{
		FieldName: AddressUnitFieldName,
		Value:     obj.AddressInput.AddressUnit,
	}
	fiedsInfo = append(fiedsInfo, field)
	//plaque
	field = fieldData{
		FieldName: AddressPlaqueFieldName,
		Value:     obj.AddressInput.AddressPlaque,
	}
	fiedsInfo = append(fiedsInfo, field)
	var updateCustomerInfo = obj.NeedUpdateCustomerInfo
	var canUpdateSelectedAddress = true
	var addNewAddress = true
	for _, field := range fiedsInfo {
		if field.ErrMsg != nil {
			updateCustomerInfo = false
			canUpdateSelectedAddress = false
			addNewAddress = false
		}
		for _, sec := range obj.Form {
			for _, inputField := range sec.Fields {
				if inputField.Name == field.FieldName {
					inputField.Value = field.Value
					inputField.Error = field.ErrMsg
				}
			}
		}
	}
	obj.UpdateCustomerInfo = updateCustomerInfo
	obj.AddNewAddress = addNewAddress
	obj.UpdateSelectedAddress = canUpdateSelectedAddress
	if va.NextChain != nil {
		newObj := va.NextChain.Next(obj)
		return newObj
	}
	return obj
}
