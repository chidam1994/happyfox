package service

import (
	"net/http"
	"time"

	"github.com/chidam1994/happyfox/contact"
	"github.com/chidam1994/happyfox/models"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type contactService struct {
	repo contact.Repository
}

func NewService(r contact.Repository) contact.Service {
	return &contactService{
		repo: r,
	}
}

func (svc *contactService) SaveContact(contact *models.Contact) (contactId uuid.UUID, err error) {
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

func (svc *contactService) FindContacts(filterMap map[models.Filter]string) ([]models.Contact, error) {
	return svc.repo.Find(filterMap)
}

func (svc *contactService) GetContact(contactId uuid.UUID) (*models.Contact, error) {
	return svc.repo.FindById(contactId)
}

func (svc *contactService) DeleteContact(contactId uuid.UUID) error {
	contact, err := svc.repo.FindById(contactId)
	if err != nil {
		return err
	}
	if contact == nil {
		return utils.GetAppError(errors.New("The contact you're trying to delete doesnt exist"), "Unable to Delete contact", http.StatusConflict)
	}
	return svc.repo.Delete(contactId)
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
