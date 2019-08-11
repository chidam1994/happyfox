package service

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
	workTag, err := models.GetTag("work")
	assert.Nil(t, err)
	personalTag, err := models.GetTag("work")
	assert.Nil(t, err)
	emails := []*models.Email{&models.Email{Id: "test@abc.com", Tag: workTag}, &models.Email{Id: "test123@abc.com", Tag: personalTag}}
	phNums := []*models.PhNum{&models.PhNum{Number: "666777", Tag: workTag}, &models.PhNum{Number: "09342", Tag: personalTag}}
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

func TestFind(t *testing.T) {
	repo := NewInMemRepo()
	svc := NewService(repo)
	workTag, err := models.GetTag("work")
	assert.Nil(t, err)
	personalTag, err := models.GetTag("work")
	assert.Nil(t, err)

	emails := []*models.Email{&models.Email{Id: "test@abc.com", Tag: workTag}, &models.Email{Id: "test123@abc.com", Tag: personalTag}}
	contact1 := &models.Contact{
		Name:   "contactOne",
		Emails: emails,
		PhNums: []*models.PhNum{},
	}
	phNums := []*models.PhNum{&models.PhNum{Number: "666777", Tag: workTag}, &models.PhNum{Number: "09342", Tag: personalTag}}
	contact2 := &models.Contact{
		Name:   "contactTwo",
		Emails: []*models.Email{},
		PhNums: phNums,
	}
	_, err = svc.SaveContact(contact1)
	assert.Nil(t, err)
	_, err = svc.SaveContact(contact2)
	assert.Nil(t, err)
	filterMap := map[Filter]string{
		NameFilter:  "one",
		PhoneFilter: "677",
	}
	results, err := svc.FindContacts(filterMap)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(results))
}
