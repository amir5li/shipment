package address

import (
	"context"
	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type checkCustomerInfo struct {
	IsValid bool
	FirstName string
	LastName string
	NationalCode string
}

type CheckCustomerEssentialInfo struct {
	NextChain AddressChain
}

func _checkInitialInfo(phone string) checkCustomerInfo {
	var essentialInfo struct {
		FirstName string `bson:"firstName"`
		LastName string `bson:"lastName"`
		NationalCode string `bson:"nationalCode"`
	}
	findRes := connection.Customer.FindOne(
		context.TODO(), 
		bson.M{"phone": phone},
		options.FindOne().SetProjection(
			bson.M{
				"nationalCode": 1,
				"firstName": 1,
				"lastName": 1,
			},
		),
	)
	findRes.Decode(&essentialInfo)
	var isValid bool = true
	if essentialInfo.FirstName == string("") || essentialInfo.LastName == string("") || essentialInfo.NationalCode == string(""){
		isValid = false
	}
	return checkCustomerInfo{
		IsValid: isValid,
		FirstName: essentialInfo.FirstName,
		LastName: essentialInfo.LastName,
		NationalCode: essentialInfo.NationalCode,
	}
}

func (cei CheckCustomerEssentialInfo) Next(obj *AddressObj) *AddressObj {
	customerInfoValidation := _checkInitialInfo(obj.UserPhone)
	var updatedForm []*AddressSection
	if customerInfoValidation.IsValid {
		for _, section := range obj.Form {
			if section.Title !=CustomerSectionTitle {
				updatedForm = append(updatedForm, section)
			}
		}
	}else{
		for _, section := range obj.Form {
			if section.Title == CustomerSectionTitle {
				for _, field := range section.Fields {
					switch field.Name{
					case CustomerFirstNameFieldName:
						field.Value = customerInfoValidation.FirstName
					case CustomerLastNameFieldName:
						field.Value = customerInfoValidation.LastName
					case CustomerNationalCodeFieldName:
						field.Value = customerInfoValidation.NationalCode
					}
				}
			}
			updatedForm = append(updatedForm, section)
		}
	}
	obj.Form = updatedForm
	if cei.NextChain != nil {
		newObj := cei.NextChain.Next(obj)
		return newObj
	}
	return obj
}