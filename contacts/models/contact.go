package models

import (
	"errors"
)

var (
	ErrEmptyNameField  = errors.New("name cannot be empty")
	ErrEmptyEmailField = errors.New("email cannot be empty")
)

// const (
// 	ERREmptyNameField  = "name cannot be empty"
// 	ERREmptyEmailField = "email cannot be empty"
// )

type Contact struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MoreInfo     string `json:"moreInfo"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	ContactNo    string `json:"contactNo"`
	Status       string `json:"status"`
	LastModified string `json:"lastModified"`
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
	return nil
}
