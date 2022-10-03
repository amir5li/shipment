package address

import (
	"context"
	"regexp"

	"github.com/amir5li/shipment/connection"
	"go.mongodb.org/mongo-driver/bson"
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
		return "0" + phone[startPhoneInd:len(phone)], nil
	}else{
		return string(""), InvalidPhoneNumber
	}
}

func _validateRawText(text string)  error{
	if match, _ := regexp.MatchString(`^[\u0600-\u06FF]{2,}[\u0600-\u06FF\s]*$`,text);match {
		return nil
	}else{
		return InvalidText
	}
}

func _validateNationalCodeCustomer(nc string) error {
	if size := len(nc); size != 10 {
		return InvalidNationalCode
	}
	count, _ := connection.Customer.CountDocuments(context.TODO(), bson.M{"nationalCode": nc})
	if count != 0 {
		return DuplicateNationalCode
	}
	return nil
}