package contactsvc

import (
	"net/http"
	"time"

	"github.com/chidam1994/happyfox/models"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (svc *Service) SaveContact(contact *models.Contact) (contactId uuid.UUID, err error) {
	contact.Id = uuid.New()
	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()
	beforeSave(contact, contact.Id)
	existingContact, err := svc.repo.FindByName(contact.Name)
	if err != nil {
		return
	}
	if existingContact != nil {
		return contactId, utils.GetAppError(errors.New("Contact with the specified name already exists"), "Unable to create contact", http.StatusConflict)
	}
	return svc.repo.Save(contact)
}

func beforeSave(contact *models.Contact, contactId uuid.UUID) {
	now := time.Now()
	for i := range contact.Emails {
		contact.Emails[i].ContactId = contactId
		contact.Emails[i].CreatedAt = now
		contact.Emails[i].UpdatedAt = now
	}
	for i := range contact.PhNums {
		contact.PhNums[i].ContactId = contactId
		contact.PhNums[i].CreatedAt = now
		contact.PhNums[i].UpdatedAt = now
	}
}
