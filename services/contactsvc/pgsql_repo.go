package contactsvc

import (
	"database/sql"
	"net/http"

	"github.com/chidam1994/happyfox/models"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
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
	trans, err := repo.DbMap.Begin()
	if err != nil {
		return contactId, utils.GetAppError(err, "error while saving contact", http.StatusInternalServerError)
	}
	err = trans.Insert(contact)
	if err != nil {
		return contactId, utils.GetAppError(err, "error while saving contact", http.StatusInternalServerError)
	}
	temp := make([]interface{}, len(contact.Emails))
	for i, val := range contact.Emails {
		temp[i] = val
	}
	err = trans.Insert(temp...)
	if err != nil {
		return contactId, utils.GetAppError(err, "error while saving contact", http.StatusInternalServerError)
	}
	temp = make([]interface{}, len(contact.PhNums))
	for i, val := range contact.PhNums {
		temp[i] = val
	}
	err = trans.Insert(temp...)
	if err != nil {
		return contactId, utils.GetAppError(err, "error while saving contact", http.StatusInternalServerError)
	}
	err = trans.Commit()
	if err != nil {
		return contactId, utils.GetAppError(err, "error while saving contact", http.StatusInternalServerError)
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
	_, err := repo.DbMap.Delete(&models.Contact{Id: contactId})
	if err != nil {
		return utils.GetAppError(err, "error while deleting contact", http.StatusInternalServerError)
	}
	return nil
}

func (repo *PgsqlRepo) Find(filterMap map[string]string) ([]*models.Contact, error) {
	panic("not implemented")
}

func (repo *PgsqlRepo) FindById(contactId uuid.UUID) (*models.Contact, error) {
	result := models.Contact{}
	err := repo.DbMap.SelectOne(&result, "select * from contacts where id= $1", contactId)
	if err != nil && err != sql.ErrNoRows {
		return nil, utils.GetAppError(err, "error while finding contact by Id", http.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &result, nil
}

func (repo *PgsqlRepo) FindByName(name string) (*models.Contact, error) {
	result := models.Contact{}
	err := repo.DbMap.SelectOne(&result, "select * from contacts where name= $1", name)
	if err != nil && err != sql.ErrNoRows {
		return nil, utils.GetAppError(err, "error while finding contact by name", http.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &result, nil
}

func (repo *PgsqlRepo) FindEmail(contactId uuid.UUID, email string) (*models.Email, error) {
	panic("not implemented")
}

func (repo *PgsqlRepo) FindPhNum(contactId uuid.UUID, phNum string) (*models.PhNum, error) {
	panic("not implemented")
}
