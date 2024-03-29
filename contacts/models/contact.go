package models

import (
	"encoding/json"
	"errors"
)

var (
	ErrEmptyNameField    = errors.New("name cannot be empty")
	ErrEmptyEmailField   = errors.New("email cannot be empty")
	ErrEmptyAddressField = errors.New("address cannot be empty")
)

// TDD

// const (
// 	ERREmptyNameField  = "name cannot be empty"
// 	ERREmptyEmailField = "email cannot be empty"
// )

type Contact struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MoreInfo     string `json:"moreInfo" gorm:"column:moreInfo"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	ContactNo    string `json:"contactNo"`
	Status       string `json:"status"`
	LastModified string `json:"lastModified" gorm:"column:lastModified"`
}

func (c *Contact) Validate() error {
	if c.Name == "" {
		//return fmt.Errorf("name filed cannot be empty")
		//return errors.New(ERREmptyNameField)
		return ErrEmptyNameField

	}
	if c.Email == "" {
		//return fmt.Errorf("email filed cannot be empty")
		//return errors.New(ERREmptyEmailField)
		return ErrEmptyEmailField
	}

	if c.Address == "" {
		return ErrEmptyAddressField
	}
	return nil
}

func (c *Contact) ToJson() (string, error) {
	buf, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
