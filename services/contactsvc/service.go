package contactsvc

import (
	"fmt"
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

type Filter string

const (
	NameFilter  Filter = "name"
	EmailFilter Filter = "email"
	PhoneFilter Filter = "phnum"
)

func GetFilter(filterStr string) (Filter, error) {
	filterMap := map[string]Filter{
		"name":  NameFilter,
		"email": EmailFilter,
		"phnum": PhoneFilter,
	}
	filter, ok := filterMap[filterStr]
	if !ok {
		return Filter(""), fmt.Errorf("error converting string %s to filter", filterStr)
	}
	return filter, nil
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

func (svc *Service) FindContacts(filterMap map[Filter]string) ([]models.Contact, error) {
	return svc.repo.Find(filterMap)
}

func (svc *Service) GetContact(contactId uuid.UUID) (*models.Contact, error) {
	return svc.repo.FindById(contactId)
}

func (svc *Service) DeleteContact(contactId uuid.UUID) error {
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
