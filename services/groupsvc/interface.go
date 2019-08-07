package groupsvc

import (
	"github.com/chidam1994/happyfox/models"
	"github.com/google/uuid"
)

type Repository interface {
	Save(group *models.Group) (uuid.UUID, error)
	AddMembers(memberIds []uuid.UUID) error
	RemMembers(memberIds []uuid.UUID) error
	GetMembersCount(memberIds []uuid.UUID) (int, error)
	RenameGroup(groupId uuid.UUID, name string) error
	Delete(groupId uuid.UUID) error
	FindByName(name string) (*models.Group, error)
	FindById(groupId uuid.UUID) (*models.Group, error)
}
