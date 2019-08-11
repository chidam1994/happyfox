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
