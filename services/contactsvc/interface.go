package contactsvc

import (
	"github.com/chidam1994/happyfox/models"
	"github.com/google/uuid"
)

type Repository interface {
	Save(contact *models.Contact) (uuid.UUID, error)
	AddEmail(contactId uuid.UUID, email *models.Email) error
	AddPhNum(contactId uuid.UUID, phNum *models.PhNum) error
	RemEmail(contactId uuid.UUID, email string) error
	RemPhNum(contactId uuid.UUID, phNum string) error
	Delete(contactId uuid.UUID) error
	Find(filterMap map[Filter]string) ([]models.Contact, error)
	FindById(contactId uuid.UUID) (*models.Contact, error)
	FindByName(name string) (*models.Contact, error)
	FindEmail(contactId uuid.UUID, email string) (*models.Email, error)
	FindPhNum(contactId uuid.UUID, phNum string) (*models.PhNum, error)
}
