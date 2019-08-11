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
	group := &models.Group{
		Name:    "testGroup",
		Members: []*models.Member{&models.Member{MemberId: uuid.New()}},
	}
	var emptyUUID uuid.UUID
	id, err := svc.SaveGroup(group)
	assert.Nil(t, err)
	assert.NotEqual(t, emptyUUID, id)
	id, err = svc.SaveGroup(group)
	assert.Equal(t, emptyUUID, id)
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	repo := NewInMemRepo()
	svc := NewService(repo)
	group := &models.Group{
		Name:    "testGroup",
		Members: []*models.Member{&models.Member{MemberId: uuid.New()}},
	}
	id, err := svc.SaveGroup(group)
	assert.Nil(t, err)
	err = svc.DeleteGroup(id)
	assert.Nil(t, err)
	err = svc.DeleteGroup(id)
	assert.NotNil(t, err)
}
