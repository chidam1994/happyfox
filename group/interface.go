package group

import (
	"github.com/chidam1994/happyfox/models"
	"github.com/google/uuid"
)

type Repository interface {
	Save(group *models.Group) (uuid.UUID, error)
	AddMembers(groupId uuid.UUID, members []*models.Member) error
	RemMembers(groupId uuid.UUID, memberIds []uuid.UUID) error
	GetMembersCount(groupId uuid.UUID, memberIds []uuid.UUID) (int, error)
	RenameGroup(groupId uuid.UUID, name string) error
	Delete(groupId uuid.UUID) error
	FindByName(name string) (*models.Group, error)
	FindById(groupId uuid.UUID) (*models.Group, error)
}

type Service interface {
	SaveGroup(group *models.Group) (uuid.UUID, error)
	DeleteGroup(groupId uuid.UUID) error
	GetGroup(groupId uuid.UUID) (*models.Group, error)
	AddMembers(groupId uuid.UUID, memberIds []uuid.UUID) error
	RemMembers(groupId uuid.UUID, memberIds []uuid.UUID) error
}
