package contactsvc

import (
	"net/http"

	"github.com/chidam1994/happyfox/models"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/gorp.v2"
)

type PgsqlRepo struct {
	DbMap *gorp.DbMap
}

func NewPgsqlRepo(dbMap *gorp.DbMap) *PgsqlRepo {
	return &PgsqlRepo{
		DbMap: dbMap,
	}
}

func (repo *PgsqlRepo) Save(contact *models.Contact) (contactId uuid.UUID, err error) {
	err = repo.DbMap.Insert(contact)
	if err != nil {
		appError := &utils.AppError{
			Code: http.StatusInternalServerError,
			Err:  errors.Wrap(err, "error while saving contact"),
		}
		return contactId, appError
	}
	return contact.Id, nil
}

func (repo *PgsqlRepo) AddEmail(contactId uuid.UUID, email *models.Email) error {
	panic("not implemented")
}

func (repo *PgsqlRepo) AddPhNum(contactId uuid.UUID, phNum *models.PhNum) error {
	panic("not implemented")
}

func (repo *PgsqlRepo) RemEmail(contactId uuid.UUID, email string) error {
	panic("not implemented")
}

func (repo *PgsqlRepo) RemPhNum(contactId uuid.UUID, phNum string) error {
	panic("not implemented")
}

func (repo *PgsqlRepo) Delete(contactId uuid.UUID) error {
	panic("not implemented")
}

func (repo *PgsqlRepo) Find(filterMap map[string]string) ([]*models.Contact, error) {
	panic("not implemented")
}

func (repo *PgsqlRepo) FindById(contactId uuid.UUID) (*models.Contact, error) {
	panic("not implemented")
}

func (repo *PgsqlRepo) FindByName(name string) (*models.Contact, error) {
	result := models.Contact{}
	err := repo.DbMap.SelectOne(&result, "select * from contacts where name= $1", name)
	if err != nil {
		appError := &utils.AppError{
			Code: http.StatusInternalServerError,
			Err:  errors.Wrap(err, "error while finding contact by name"),
		}
		return nil, appError
	}
	return &result, nil
}

func (repo *PgsqlRepo) FindEmail(contactId uuid.UUID, email string) (*models.Email, error) {
	panic("not implemented")
}

func (repo *PgsqlRepo) FindPhNum(contactId uuid.UUID, phNum string) (*models.PhNum, error) {
	panic("not implemented")
}
