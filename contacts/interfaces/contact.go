package interfaces

import "contacts/models"

type IContact interface {
	Create(*models.Contact) (interface{}, error)
	GetBy(id string) (*models.Contact, error)
	UpdateBy(id string, data map[string]interface{}) (interface{}, error)
	DeleteBy(id string) (interface{}, error)
}
