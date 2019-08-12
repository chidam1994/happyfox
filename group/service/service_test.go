package service

import (
	"testing"

	"github.com/chidam1994/happyfox/group/repository"
	"github.com/chidam1994/happyfox/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	repo := repository.NewInMemRepo()
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
	repo := repository.NewInMemRepo()
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

func TestAddMembers(t *testing.T) {
	repo := repository.NewInMemRepo()
	svc := NewService(repo)
	memId1 := uuid.New()
	memId2 := uuid.New()
	memId3 := uuid.New()
	group := &models.Group{
		Name:    "testGroup",
		Members: []*models.Member{&models.Member{MemberId: memId1}, &models.Member{MemberId: memId2}},
	}
	id, err := svc.SaveGroup(group)
	assert.Nil(t, err)
	err = svc.AddMembers(id, []uuid.UUID{memId3, uuid.New()})
	assert.NoError(t, err)
	err = svc.AddMembers(id, []uuid.UUID{memId2})
	assert.Error(t, err)
}

func TestRemMembers(t *testing.T) {
	repo := repository.NewInMemRepo()
	svc := NewService(repo)
	memId1 := uuid.New()
	memId2 := uuid.New()
	group := &models.Group{
		Name:    "testGroup",
		Members: []*models.Member{&models.Member{MemberId: memId1}, &models.Member{MemberId: memId2}},
	}
	id, err := svc.SaveGroup(group)
	assert.Nil(t, err)
	err = svc.RemMembers(id, []uuid.UUID{memId1, memId2})
	assert.NoError(t, err)
	err = svc.RemMembers(id, []uuid.UUID{memId2})
	assert.Error(t, err)
}
