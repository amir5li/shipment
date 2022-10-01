package address

import (
	"context"
	"github.com/amir5li/shipment/location"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GenRawForm struct {
	NextChain AddressChain
}

func (grf GenRawForm) Next(obj *AddressObj) *AddressObj {
	var rawForm []*AddressSection
	var customerSection *AddressSection = &AddressSection{}
	customerSection.Title = CustomerSectionTitle
	var customerSectionFields []*AddressField
	customerSectionFields = append(customerSectionFields,
		&AddressField{
			Name:     CustomerFirstNameFieldName,
			Label:    CustomerLastNameLabelName,
			Type:     "text",
			Required: true,
			Value:    nil,
		},
		&AddressField{
			Name:     CustomerLastNameFieldName,
			Label:    CustomerLastNameLabelName,
			Type:     "text",
			Required: true,
			Value:   nil,
		},
		&AddressField{
			Name:     CustomerNationalCodeFieldName,
			Label:    CustomerNationalCodeLabelName,
			Type:     "text",
			Required: true,
			Value:    nil,
		},
	)
	customerSection.Fields = customerSectionFields
	rawForm = append(rawForm, customerSection)
	var consigneeSection *AddressSection = &AddressSection{}
	consigneeSection.Title = ConsigneeSectionTitle
	var consigneeSectionFields []*AddressField
	consigneeSectionFields = append(consigneeSectionFields,
		&AddressField{
			Name:     ConsigneeIsCustomerFieldName,
			Label:    ConsigneeIsCustomerLabelName,
			Required: true,
			Type:     "checkbox",
			Value:    true,
		},
		&AddressField{
			Name:     ConsigneeFirstNameFieldName,
			Label:    ConsigneeFirstNameLabelName,
			Type:     "text",
			Required: false,
			Value:    nil,
		},
		&AddressField{
			Name:     ConsigneeLastNameFieldName,
			Label:    ConsigneeLastNameLabelName,
			Type:     "text",
			Required: false,
			Value:    nil,
		},
		&AddressField{
			Name:     ConsigneeNationalCodeFieldName,
			Label:    ConsigneeNationalCodeLabelName,
			Type:     "text",
			Required: false,
			Value:    nil,
		},
		&AddressField{
			Name:     ConsigneePhoneFieldName,
			Label:    ConsigneePhoneLabelName,
			Type:     "text",
			Required: false,
			Value:    nil,
		},
	)
	consigneeSection.Fields = consigneeSectionFields
	rawForm = append(rawForm, consigneeSection)
	var addressSection *AddressSection = &AddressSection{}
	addressSection.Title = AddressSectionTitle
	var addressSectionFields []*AddressField
	addressSectionFields = append(addressSectionFields,
		&AddressField{
			Name:     AddressCityFieldName,
			Label:    AddressCityLabelName,
			Type:     "select",
			Required: true,
			Value:    primitive.NilObjectID,
			Options:  []location.Loc{{"--", primitive.NilObjectID}},
		},
		&AddressField{
			Name:     AddressProvinceFieldName,
			Label:    AddressProvinceLabelName,
			Required: true,
			Type:     "select",
			Value:    primitive.NilObjectID,
			Options:  location.GetProvinceList(context.TODO()),
		},
		&AddressField{
			Name:     AddressPostalAddressFieldName,
			Label:    AddressPostalAddressLabelName,
			Required: true,
			Type:     "text",
			Value:    nil,
		},
		&AddressField{
			Name:     AddressPostalCodeFieldName,
			Label:    AddressPostalCodeLabelName,
			Required: false,
			Type:     "number",
			Value:    nil,
		},
		&AddressField{
			Name:     AddressUnitFieldName,
			Label:    AddressUnitLabelName,
			Required: false,
			Type:     "number",
			Value:    nil,
		},
		&AddressField{
			Name:     AddressPlaqueFieldName,
			Label:    AddressPlaqueLabelName,
			Required: true,
			Type:     "number",
			Value:    nil,
		},
	)
	addressSection.Fields = addressSectionFields
	rawForm = append(rawForm, addressSection)
	obj.Form = rawForm
	if grf.NextChain != nil {
		newObj := grf.NextChain.Next(obj)
		return newObj
	}
	return obj
}
