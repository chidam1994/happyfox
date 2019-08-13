package repository

import (
	"database/sql"
	"fmt"
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

func (repo *PgsqlRepo) Delete(contactId uuid.UUID) error {
	_, err := repo.DbMap.Delete(&models.Contact{Id: contactId})
	if err != nil {
		return utils.GetAppError(err, "error while deleting contact", http.StatusInternalServerError)
	}
	return nil
}

func (repo *PgsqlRepo) Find(filterMap map[models.Filter]string) ([]models.Contact, error) {
	results := []models.Contact{}
	var query string
	if len(filterMap) > 0 {
		query = fmt.Sprintf("select distinct contacts.id, contacts.name, contacts.created_at from contacts left join emails on contacts.id = emails.contact_id left join phnumbers on contacts.id = phnumbers.contact_id where %s order by contacts.created_at", GetSearchCondition(filterMap))
	} else {
		return results, utils.GetAppError(errors.New("no search filters specified"), "error while searching for contact", http.StatusBadRequest)
	}
	_, err := repo.DbMap.Select(&results, query)
	if err != nil {
		return results, utils.GetAppError(err, "error while deleting contact", http.StatusInternalServerError)
	}
	return results, nil
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
	emails := []*models.Email{}
	_, err = repo.DbMap.Select(&emails, "select * from emails where contact_id= $1", contactId)
	if err != nil && err != sql.ErrNoRows {
		return nil, utils.GetAppError(err, "error while finding contact by Id", http.StatusInternalServerError)
	}
	result.Emails = emails
	phNums := []*models.PhNum{}
	_, err = repo.DbMap.Select(&phNums, "select * from phnumbers where contact_id= $1", contactId)
	if err != nil && err != sql.ErrNoRows {
		return nil, utils.GetAppError(err, "error while finding contact by Id", http.StatusInternalServerError)
	}
	result.PhNums = phNums
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

func GetSearchCondition(filtersMap map[models.Filter]string) string {
	result := ""
	if value, ok := filtersMap[models.NameFilter]; ok {
		result = result + fmt.Sprintf("contacts.name like '%%%s%%' ", value)
	}
	if value, ok := filtersMap[models.EmailFilter]; ok {
		if len(result) > 0 {
			result = result + "or "
		}
		result = result + fmt.Sprintf("emails.email_id like '%%%s%%' ", value)
	}
	if value, ok := filtersMap[models.PhoneFilter]; ok {
		if len(result) > 0 {
			result = result + "or "
		}
		result = result + fmt.Sprintf("phnumbers.phnum like '%%%s%%'", value)
	}
	return result
}
