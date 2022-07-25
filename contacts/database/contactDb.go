package database

import (
	"contacts/models"
	"strconv"

	"gorm.io/gorm"
)

type ContactDB struct {
	DB interface{} // database connection type can be any
}

func (c *ContactDB) Create(contact *models.Contact) (*models.Contact, error) {

	if err := c.DB.(*gorm.DB).AutoMigrate(&models.Contact{}); err != nil {
		return nil, err
	}

	// Create
	tx := c.DB.(*gorm.DB).Create(contact)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return contact, nil
}
func (c *ContactDB) GetBy(id string) (*models.Contact, error) {
	if id == "" {
		return nil, ErrInvalidID
	}
	contact := &models.Contact{}
	tx := c.DB.(*gorm.DB).First(contact, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return contact, nil
}

func (c *ContactDB) UpdateBy(id string, data map[string]interface{}) (*models.Contact, error) {
	//db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, ErrInvalidID
	}
	contact := &models.Contact{ID: _id}
	tx := c.DB.(*gorm.DB).Model(contact).Updates(data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	contact, err = c.GetBy(id)
	if err != nil {
		return nil, tx.Error
	}
	return contact, nil
}

func (c *ContactDB) DeleteBy(id string) (interface{}, error) {
	tx := c.DB.(*gorm.DB).Delete(&models.Contact{}, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx.RowsAffected, nil
}
