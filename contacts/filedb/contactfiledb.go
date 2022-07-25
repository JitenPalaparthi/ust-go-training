package filedb

import (
	"contacts/models"
	"os"
)

type ContactFileDB struct {
	File string
}

func (c *ContactFileDB) Create(contact *models.Contact) (*models.Contact, error) {
	str, err := contact.ToJson()
	if err != nil {
		return nil, err
	}

	f, err := os.Create(c.File) // 666 Read|Write Read|Write Read|Write
	if err != nil {
		return nil, err
	}
	_, err = f.WriteString(str)

	if err != nil {
		return nil, err
	}

	return contact, nil
}
