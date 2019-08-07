package contactsvc

import (
	"testing"

	"github.com/chidam1994/happyfox/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	repo := NewInMemRepo()
	svc := NewService(repo)
	contact := &models.Contact{
		Name:   "testContact",
		Emails: []*models.Email{},
		PhNums: []*models.PhNum{},
	}
	var emptyUUID uuid.UUID
	id, err := svc.SaveContact(contact)
	assert.Nil(t, err)
	assert.NotEqual(t, emptyUUID, id)
	id, err = svc.SaveContact(contact)
	assert.Equal(t, emptyUUID, id)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	repo := NewInMemRepo()
	svc := NewService(repo)
	emails := []*models.Email{&models.Email{Id: "test@abc.com", Tag: models.GetTag("work")}, &models.Email{Id: "test123@abc.com", Tag: models.GetTag("personal")}}
	phNums := []*models.PhNum{&models.PhNum{Number: "666777", Tag: models.GetTag("work")}, &models.PhNum{Number: "09342", Tag: models.GetTag("personal")}}
	contact := &models.Contact{
		Name:   "testContact",
		Emails: emails,
		PhNums: phNums,
	}
	id, err := svc.SaveContact(contact)
	assert.Nil(t, err)
	err = svc.DeleteContact(id)
	assert.Nil(t, err)
	err = svc.DeleteContact(id)
	assert.NotNil(t, err)
}
