package groupsvc

import (
	"errors"

	"github.com/chidam1994/happyfox/models"
	"github.com/google/uuid"
)

type InMemRepo struct {
	groupsMap map[uuid.UUID]*models.Group
}

func NewInMemRepo() *InMemRepo {
	return &InMemRepo{
		groupsMap: make(map[uuid.UUID]*models.Group, 5),
	}
}

func (repo *InMemRepo) Save(group *models.Group) (uuid.UUID, error) {
	repo.groupsMap[group.Id] = group
	return group.Id, nil

}

func (repo *InMemRepo) AddMembers(memberIds []uuid.UUID) error {
	panic("not implemented")
}

func (repo *InMemRepo) RemMembers(memberIds []uuid.UUID) error {
	panic("not implemented")
}

func (repo *InMemRepo) GetMembersCount(memberIds []uuid.UUID) (int, error) {
	panic("not implemented")
}

func (repo *InMemRepo) RenameGroup(groupId uuid.UUID, name string) error {
	panic("not implemented")
}

func (repo *InMemRepo) Delete(groupId uuid.UUID) error {
	panic("not implemented")
}

func (repo *InMemRepo) FindByName(name string) (*models.Group, error) {
	for _, group := range repo.groupsMap {
		if name == group.Name {
			return group, nil
		}
	}
	return nil, nil
}

func (repo *InMemRepo) FindById(groupId uuid.UUID) (*models.Group, error) {
	group, ok := repo.groupsMap[groupId]
	if !ok {
		return nil, errors.New("Contact Not found")
	}
	return group, nil
}
