package groupsvc

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
		Members: []uuid.UUID{uuid.New(), uuid.New(), uuid.New()},
	}
	var emptyUUID uuid.UUID
	id, err := svc.SaveGroup(group)
	assert.Nil(t, err)
	assert.NotEqual(t, emptyUUID, id)
	id, err = svc.SaveGroup(group)
	assert.Equal(t, emptyUUID, id)
	assert.NotNil(t, err)
}
