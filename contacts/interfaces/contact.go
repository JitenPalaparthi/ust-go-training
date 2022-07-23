package interfaces

import "contacts/models"

type IContact interface {
	Create(*models.Contact) (*models.Contact, error)
	GetBy(id string) (*models.Contact, error)
	UpdateBy(id string, data map[string]interface{}) (*models.Contact, error)
	DeleteBy(id string) (interface{}, error)
}
