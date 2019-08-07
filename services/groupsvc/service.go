package groupsvc

import (
	"errors"
	"net/http"
	"time"

	"github.com/chidam1994/happyfox/models"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (svc *Service) SaveGroup(group *models.Group) (groupId uuid.UUID, err error) {
	group.Id = uuid.New()
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()
	existingGroup, err := svc.repo.FindByName(group.Name)
	if err != nil {
		return groupId, err
	}
	if existingGroup != nil {
		return groupId, utils.GetAppError(errors.New("Group with the specified name already exists"), "Unable to create group", http.StatusConflict)
	}
	return svc.repo.Save(group)
}
