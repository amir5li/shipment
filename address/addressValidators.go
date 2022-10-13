package address

import (
	"context"
	"regexp"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func _validatePhone(phone string) (string, error){
	if match, _ :=regexp.MatchString(`^09\d{9}$`, phone); match{
		return phone, nil
	}else if match, _ := regexp.MatchString(`9\d{9}`, phone); match {
		return "0" + phone, nil
	}else if match, _ := regexp.MatchString(`^((0)|(+))?989\d{9}$`, phone); match {
		reg, _ := regexp.Compile(`^((0)|(+))?98`)
		matchedInds  := reg.FindIndex([]byte(phone))
		startPhoneInd := matchedInds[1]
		return "0" + phone[startPhoneInd:], nil
	}else{
		return string(""), InvalidPhoneNumber
	}
}

func _validateRawText(text string)  error{
	if match, _ := regexp.MatchString(`^[\w]{2,}[\s\w]*$`,text);match {
		return nil
	}else{
		return InvalidText
	}
}

func _validateNationalCodeCustomer(nc string) error {
	match, _ := regexp.MatchString(`^\d{8,10}$`,nc)
	if !match {
		return InvalidNationalCode
	}
	count, _ := connection.Customer.CountDocuments(context.TODO(), bson.M{"nationalCode": nc})
	if count != 0 {
		return DuplicateNationalCode
	}
	return nil
}

func _validateNationalCodeConsignee(nc string) error {
	if match, _ := regexp.MatchString(`^\d{8,10}$`, nc); !match {
		return InvalidNationalCode
	}
	return nil
}

func _validateAddressPostal(addr string) error {
	match, _ := regexp.MatchString(`[^\w\d\s\(\)\,]`, addr)
	if match || len(addr) < 10 {
		return InvalidPostalAddress
	}
	return nil
}

func _validateAddressPostalCode(code string) error {
	match, _ := regexp.MatchString(`^\d{10}$`, code)
	if !match{
		return InvalidPostalCode
	}
	return nil
}
func _validateAddressProvince(provinceID primitive.ObjectID) error {
	count, _ := connection.Province.CountDocuments(context.TODO(), bson.M{"_id": provinceID})
	if count != 1 {
		return InvalidProvince
	}
	return nil
}

func _validateAddressCity(cityID primitive.ObjectID, provinceID primitive.ObjectID) error {
	count, _ := connection.City.CountDocuments(context.TODO(), bson.M{"_id": cityID, "provinceID": provinceID})
	if count != 1{
		return InvalidCity
	}
	return nil
}