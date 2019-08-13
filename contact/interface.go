package contact

import (
	"github.com/chidam1994/happyfox/models"
	"github.com/google/uuid"
)

type Repository interface {
	Save(contact *models.Contact) (uuid.UUID, error)
	Delete(contactId uuid.UUID) error
	Find(filterMap map[models.Filter]string) ([]models.Contact, error)
	FindById(contactId uuid.UUID) (*models.Contact, error)
	FindByName(name string) (*models.Contact, error)
}

type Service interface {
	SaveContact(contact *models.Contact) (uuid.UUID, error)
	FindContacts(filterMap map[models.Filter]string) ([]models.Contact, error)
	DeleteContact(contactId uuid.UUID) error
	GetContact(contactId uuid.UUID) (*models.Contact, error)
}
