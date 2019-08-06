package contactsvc

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

func (repo *InMemRepo) AddEmail(contactId uuid.UUID, email *models.Email) (err error) {
	contact, ok := repo.contactsMap[contactId]
	if !ok {
		return errors.New("Contact Not found")
	}
	contact.Emails = append(contact.Emails, email)
	return
}

func (repo *InMemRepo) RemEmail(contactId uuid.UUID, email string) (err error) {
	contact, ok := repo.contactsMap[contactId]
	if !ok {
		return errors.New("Contact Not found")
	}
	for i := range contact.Emails {
		if contact.Emails[i].Id == email {
			contact.Emails = append(contact.Emails[:i], contact.Emails[i+1])
			break
		}
	}
	return
}

func (repo *InMemRepo) AddPhNum(contactId uuid.UUID, phNum *models.PhNum) (err error) {
	contact, ok := repo.contactsMap[contactId]
	if !ok {
		return errors.New("Contact Not found")
	}
	contact.PhNums = append(contact.PhNums, phNum)
	return
}

func (repo *InMemRepo) RemPhNum(contactId uuid.UUID, phNum string) (err error) {
	contact, ok := repo.contactsMap[contactId]
	if !ok {
		return errors.New("Contact Not found")
	}
	for i := range contact.PhNums {
		if contact.PhNums[i].Number == phNum {
			contact.PhNums = append(contact.PhNums[:i], contact.PhNums[i+1])
			break
		}
	}
	return
}

func (repo *InMemRepo) Delete(contactId uuid.UUID) (err error) {
	_, ok := repo.contactsMap[contactId]
	if !ok {
		return errors.New("Contact Not found")
	}
	delete(repo.contactsMap, contactId)
	return nil
}

func (repo *InMemRepo) Find(filterMap map[string]string) (results []*models.Contact, err error) {
	for _, contact := range repo.contactsMap {
		addToResult := true
		if filterValue, ok := filterMap["name"]; ok {
			if !strings.Contains(strings.ToLower(contact.Name), strings.ToLower(filterValue)) {
				addToResult = false
				continue
			}
		}
		if filterValue, ok := filterMap["email"]; ok {
			emailMatch := false
			for _, email := range contact.Emails {
				if strings.Contains(strings.ToLower(email.Id), strings.ToLower(filterValue)) {
					emailMatch = true
					break
				}
			}
			if !emailMatch {
				addToResult = false
				continue
			}
		}
		if filterValue, ok := filterMap["phNum"]; ok {
			phNumMatch := false
			for _, phNum := range contact.PhNums {
				if strings.Contains(strings.ToLower(phNum.Number), strings.ToLower(filterValue)) {
					phNumMatch = true
					break
				}
			}
			if !phNumMatch {
				addToResult = false
				continue
			}
		}
		if addToResult {
			results = append(results, contact)
		}
	}
	return
}

func (repo *InMemRepo) FindById(contactId uuid.UUID) (contact *models.Contact, err error) {
	contact, ok := repo.contactsMap[contactId]
	if !ok {
		return nil, errors.New("Contact Not found")
	}
	return
}

func (repo *InMemRepo) FindByName(name string) (contact *models.Contact, err error) {
	for _, contact := range repo.contactsMap {
		if name == contact.Name {
			return contact, nil
		}
	}
	return nil, errors.New("Contact Not found")
}

func (repo *InMemRepo) FindEmail(contactId uuid.UUID, emailstr string) (email *models.Email, err error) {
	contact, ok := repo.contactsMap[contactId]
	if !ok {
		return nil, errors.New("Contact Not found")
	}
	for _, email = range contact.Emails {
		if email.Id == emailstr {
			return
		}
	}
	return nil, errors.New("Email Id not found")
}

func (repo *InMemRepo) FindPhNum(contactId uuid.UUID, phNumstr string) (phNum *models.PhNum, err error) {
	contact, ok := repo.contactsMap[contactId]
	if !ok {
		return nil, errors.New("Contact Not found")
	}
	for _, phNum = range contact.PhNums {
		if phNum.Number == phNumstr {
			return
		}
	}
	return nil, errors.New("phone number not found")
}
