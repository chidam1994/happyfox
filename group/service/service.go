package service

import (
	"errors"
	"net/http"
	"time"

	"github.com/chidam1994/happyfox/group"
	"github.com/chidam1994/happyfox/models"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
)

type groupService struct {
	repo group.Repository
}

func NewService(r group.Repository) group.Service {
	return &groupService{
		repo: r,
	}
}

func (svc *groupService) SaveGroup(group *models.Group) (groupId uuid.UUID, err error) {
	group.Id = uuid.New()
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()
	beforeSave(group, group.Id)
	existingGroup, err := svc.repo.FindByName(group.Name)
	if err != nil {
		return groupId, err
	}
	if existingGroup != nil {
		return groupId, utils.GetAppError(errors.New("Group with the specified name already exists"), "Unable to create group", http.StatusConflict)
	}
	return svc.repo.Save(group)
}

func (svc *groupService) DeleteGroup(groupId uuid.UUID) error {
	group, err := svc.repo.FindById(groupId)
	if err != nil {
		return err
	}
	if group == nil {
		return utils.GetAppError(errors.New("The group you're trying to delete doesnt exist"), "Unable to Delete group", http.StatusBadRequest)
	}
	return svc.repo.Delete(groupId)
}

func (svc *groupService) GetGroup(groupId uuid.UUID) (*models.Group, error) {
	return svc.repo.FindById(groupId)
}

func (svc *groupService) AddMembers(groupId uuid.UUID, memberIds []uuid.UUID) error {
	group, err := svc.repo.FindById(groupId)
	if err != nil {
		return err
	}
	if group == nil {
		return utils.GetAppError(errors.New("The group you're trying to add members to doesnt exist"), "Unable to add members", http.StatusBadRequest)
	}
	num, err := svc.repo.GetMembersCount(groupId, memberIds)
	if err != nil {
		return err
	}
	if num > 0 {
		return utils.GetAppError(errors.New("some of the contacts you are trying to add are already members of the group"), "Unable to Add members", http.StatusBadRequest)
	}
	return svc.repo.AddMembers(groupId, getMembers(memberIds, groupId))
}

func (svc *groupService) RemMembers(groupId uuid.UUID, memberIds []uuid.UUID) error {
	group, err := svc.repo.FindById(groupId)
	if err != nil {
		return err
	}
	if group == nil {
		return utils.GetAppError(errors.New("The group you're trying to add members to doesnt exist"), "Unable to add members", http.StatusBadRequest)
	}
	num, err := svc.repo.GetMembersCount(groupId, memberIds)
	if err != nil {
		return err
	}
	if num < len(memberIds) {
		return utils.GetAppError(errors.New("some of the contacts you are trying to remove are not members of the group"), "Unable to remove members", http.StatusBadRequest)
	}
	return svc.repo.RemMembers(groupId, memberIds)
}

func beforeSave(group *models.Group, groupId uuid.UUID) {
	now := time.Now()
	for i := range group.Members {
		group.Members[i].GroupId = groupId
		group.Members[i].CreatedAt = now
		group.Members[i].UpdatedAt = now
	}
}

func getMembers(memberIds []uuid.UUID, groupId uuid.UUID) []*models.Member {
	now := time.Now()
	result := make([]*models.Member, len(memberIds))
	for i := range memberIds {
		result[i] = &models.Member{
			MemberId:  memberIds[i],
			GroupId:   groupId,
			CreatedAt: now,
			UpdatedAt: now,
		}
	}
	return result
}
