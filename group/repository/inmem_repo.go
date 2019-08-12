package repository

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

func (repo *InMemRepo) AddMembers(groupId uuid.UUID, members []*models.Member) error {
	group := repo.groupsMap[groupId]
	group.Members = append(group.Members, members...)
	return nil
}

func (repo *InMemRepo) RemMembers(groupId uuid.UUID, memberIds []uuid.UUID) error {
	group := repo.groupsMap[groupId]
	for _, memId := range memberIds {
		for i := range group.Members {
			if group.Members[i].MemberId == memId {
				group.Members = append(group.Members[:i], group.Members[i+1:]...)
				break
			}
		}
	}
	return nil
}

func (repo *InMemRepo) GetMembersCount(groupId uuid.UUID, memberIds []uuid.UUID) (int, error) {
	count := 0
	for _, id := range memberIds {
		for _, member := range repo.groupsMap[groupId].Members {
			if id == member.MemberId {
				count += 1
				continue
			}
		}
	}
	return count, nil
}

func (repo *InMemRepo) RenameGroup(groupId uuid.UUID, name string) error {
	panic("not implemented")
}

func (repo *InMemRepo) Delete(groupId uuid.UUID) error {
	delete(repo.groupsMap, groupId)
	return nil
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
