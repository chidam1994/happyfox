package transport

import (
	"errors"

	"github.com/chidam1994/happyfox/models"
	"github.com/google/uuid"
)

type GroupId struct {
	Id string `json:"id"`
}

type CreateGroupRequest struct {
	Name      string   `json:"name"`
	MemberIds []string `json:"members"`
}

type UpdateGroupRequest struct {
	GroupId   string   `json:"group_id"`
	Name      string   `json:"name"`
	MemberIds []string `json:"members"`
}

type updateAction string

const (
	addMembers  updateAction = "addmembers"
	remMembers  updateAction = "remmembers"
	renameGroup updateAction = "rename"
)

func getUpdateAction(actionStr string) (updateAction, error) {
	updateActionMap := map[string]updateAction{
		"addmembers": addMembers,
		"remmembers": remMembers,
		"rename":     renameGroup,
	}
	action, ok := updateActionMap[actionStr]
	if !ok {
		return updateAction(""), errors.New("Invalid update action")
	}
	return action, nil
}

func (request *CreateGroupRequest) Validate() (*models.Group, error) {
	result := &models.Group{}
	if request.Name == "" {
		return nil, errors.New("name cannot be blank")
	}
	result.Name = request.Name
	for i := range request.MemberIds {
		if request.MemberIds[i] == "" {
			return nil, errors.New("invalid contactIds present in members")
		}
		id, err := uuid.Parse(request.MemberIds[i])
		if err != nil {
			return nil, errors.New("invalid contactIds present in members")
		}
		result.Members = append(result.Members, &models.Member{MemberId: id})
	}
	return result, nil
}

func validateUuids(uuids []string) ([]uuid.UUID, error) {
	result := make([]uuid.UUID, len(uuids))
	for i, uuidstr := range uuids {
		uuid, err := uuid.Parse(uuidstr)
		if err != nil {
			return result, errors.New("invalid uuid")
		}
		result[i] = uuid
	}
	return result, nil
}
