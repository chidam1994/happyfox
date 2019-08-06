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
