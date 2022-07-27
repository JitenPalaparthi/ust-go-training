package models_test

import (
	"contacts/models"
	"testing"
)

func TestValidateSuccess(t *testing.T) {
	contact := new(models.Contact)
	contact.Name = "UST Global"
	contact.Email = "info@ustglobal.com"
	contact.Address = "WhiteField, Bangalore"
	if err := contact.Validate(); err != nil {
		t.Fail()
	}
}

func TestValidateFail(t *testing.T) {
	contact := new(models.Contact)
	contact.Email = "info@ustglobal.com"
	contact.Address = "WhiteField, Bangalore"
	err := contact.Validate()
	if err == nil {
		t.Fail()
	}

	if err.Error() != "name cannot be empty" {
		t.Fail()
	}
}

func TestToJonSuccess(t *testing.T) {
	contact := new(models.Contact)
	contact.ID = 101
	contact.Name = "UST Global"
	contact.MoreInfo = "Demo purpose"
	contact.Email = "info@ustglobal.com"
	contact.Address = "WhiteField, Bangalore"
	contact.ContactNo = "9618558500"
	contact.Status = "created"
	contact.LastModified = "27-07-2022"

	expected := `{"id":101,"name":"UST Global","moreInfo":"Demo purpose","address":"WhiteField, Bangalore","email":"info@ustglobal.com","contactNo":"9618558500","status":"created","lastModified":"27-07-2022"}`

	actual, err := contact.ToJson()
	t.Log(actual)
	if actual != expected || err != nil {
		t.Fail()
	}

}
