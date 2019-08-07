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
	existingContact, err := svc.repo.FindByName(contact.Name)
	if err != nil {
		return
	}
	if existingContact != nil {
		appError := &utils.AppError{
			Code: http.StatusConflict,
			Err:  errors.Wrap(errors.New("Contact with the specified Name already exists"), "Unable to create contact"),
		}
		return contactId, appError
	}
	return svc.repo.Save(contact)
}
