package repository

import (
	"strings"

	"github.com/chidam1994/happyfox/models"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type InMemRepo struct {
	contactsMap map[uuid.UUID]*models.Contact
}

func NewInMemRepo() *InMemRepo {
	return &InMemRepo{
		contactsMap: make(map[uuid.UUID]*models.Contact, 5),
	}
}

func (repo *InMemRepo) Save(contact *models.Contact) (uuid.UUID, error) {
	repo.contactsMap[contact.Id] = contact
	return contact.Id, nil
}

func (repo *InMemRepo) Delete(contactId uuid.UUID) (err error) {
	delete(repo.contactsMap, contactId)
	return nil
}

func (repo *InMemRepo) Find(filterMap map[models.Filter]string) (results []models.Contact, err error) {
	for _, contact := range repo.contactsMap {
		if filterValue, ok := filterMap[models.NameFilter]; ok {
			if strings.Contains(strings.ToLower(contact.Name), strings.ToLower(filterValue)) {
				results = append(results, *contact)
				continue
			}
		}
		if filterValue, ok := filterMap[models.EmailFilter]; ok {
			emailMatch := false
			for _, email := range contact.Emails {
				if strings.Contains(strings.ToLower(email.Id), strings.ToLower(filterValue)) {
					emailMatch = true
					break
				}
			}
			if emailMatch {
				results = append(results, *contact)
				continue
			}
		}
		if filterValue, ok := filterMap[models.PhoneFilter]; ok {
			phNumMatch := false
			for _, phNum := range contact.PhNums {
				if strings.Contains(strings.ToLower(phNum.Number), strings.ToLower(filterValue)) {
					phNumMatch = true
					break
				}
			}
			if phNumMatch {
				results = append(results, *contact)
				continue
			}
		}
	}
	return
}

func (repo *InMemRepo) FindById(contactId uuid.UUID) (contact *models.Contact, err error) {
	contact, ok := repo.contactsMap[contactId]
	if !ok {
		return nil, errors.New("Contact Not found")
	}
	return contact, nil
}

func (repo *InMemRepo) FindByName(name string) (contact *models.Contact, err error) {
	for _, contact := range repo.contactsMap {
		if name == contact.Name {
			return contact, nil
		}
	}
	return nil, nil
}
