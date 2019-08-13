package transport

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chidam1994/happyfox/models"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type mockGroupService struct {
}

func (mockSvc *mockGroupService) SaveGroup(group *models.Group) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (mockSvc *mockGroupService) DeleteGroup(groupId uuid.UUID) error {
	return nil
}

func (mockSvc *mockGroupService) GetGroup(groupId uuid.UUID) (*models.Group, error) {
	mockGroup := &models.Group{Name: "mockGroup",
		Members: []*models.Member{&models.Member{MemberId: uuid.New()}, &models.Member{MemberId: uuid.New()}}}
	return mockGroup, nil
}
func (mockSvc *mockGroupService) AddMembers(groupId uuid.UUID, memberIds []uuid.UUID) error {
	return nil
}
func (mockSvc *mockGroupService) RemMembers(groupId uuid.UUID, memberIds []uuid.UUID) error {
	return nil
}
func (mockSvc *mockGroupService) RenameGroup(groupId uuid.UUID, name string) error {
	return nil
}

func TestCreateGroup(t *testing.T) {
	group := &CreateGroupRequest{
		Name:      "testgroup",
		MemberIds: []string{uuid.New().String(), uuid.New().String()},
	}
	svc := &mockGroupService{}
	handler := &groupHandler{
		groupSvc: svc,
	}
	createGroupReqByteArr, err := json.Marshal(group)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/group", strings.NewReader(string(createGroupReqByteArr)))
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	httphandler := http.HandlerFunc(handler.createGroup)
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var respBody GroupId
	err = json.Unmarshal([]byte(rr.Body.String()), &respBody)
	assert.NoError(t, err)
	_, err = uuid.Parse(respBody.Id)
	assert.NoError(t, err)
}

func TestCreateGroupErrors(t *testing.T) {
	group1 := &CreateGroupRequest{
		MemberIds: []string{uuid.New().String(), uuid.New().String()},
	}
	group2 := &CreateGroupRequest{
		Name:      "mockgrp1",
		MemberIds: []string{"invalid UUID", uuid.New().String()},
	}
	groupReqs := []*CreateGroupRequest{group1, group2}
	svc := &mockGroupService{}
	handler := &groupHandler{
		groupSvc: svc,
	}
	for _, grp := range groupReqs {
		createGroupReqByteArr, err := json.Marshal(grp)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/group", strings.NewReader(string(createGroupReqByteArr)))
		assert.NoError(t, err)
		rr := httptest.NewRecorder()
		httphandler := http.HandlerFunc(handler.createGroup)
		httphandler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusBadRequest, rr.Code)
		var respBody utils.Response
		err = json.Unmarshal([]byte(rr.Body.String()), &respBody)
		assert.NoError(t, err)
		assert.Equal(t, respBody.Success, false)
		assert.NotNil(t, respBody.Error)
	}
}

func TestDeleteGroup(t *testing.T) {
	groupId := uuid.New().String()
	svc := &mockGroupService{}
	handler := &groupHandler{
		groupSvc: svc,
	}
	httphandler := http.HandlerFunc(handler.deleteGroup)
	req, err := http.NewRequest("DELETE", "/group/"+groupId, nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": groupId})
	rr := httptest.NewRecorder()
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestDeleteGroupErrors(t *testing.T) {
	groupId := "Invalid Id"
	svc := &mockGroupService{}
	handler := &groupHandler{
		groupSvc: svc,
	}
	httphandler := http.HandlerFunc(handler.deleteGroup)
	req, err := http.NewRequest("DELETE", "/group/"+groupId, nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": groupId})
	rr := httptest.NewRecorder()
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	var respBody utils.Response
	err = json.Unmarshal([]byte(rr.Body.String()), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, respBody.Success, false)
	assert.NotNil(t, respBody.Error)
}

func TestGetGroup(t *testing.T) {
	groupId := uuid.New().String()
	svc := &mockGroupService{}
	handler := &groupHandler{
		groupSvc: svc,
	}
	httphandler := http.HandlerFunc(handler.getGroup)
	req, err := http.NewRequest("GET", "/group/"+groupId, nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": groupId})
	rr := httptest.NewRecorder()
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetGroupErrors(t *testing.T) {
	groupId := "Invalid Id"
	svc := &mockGroupService{}
	handler := &groupHandler{
		groupSvc: svc,
	}
	httphandler := http.HandlerFunc(handler.getGroup)
	req, err := http.NewRequest("GET", "/group/"+groupId, nil)
	assert.NoError(t, err)
	req = mux.SetURLVars(req, map[string]string{"id": groupId})
	rr := httptest.NewRecorder()
	httphandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	var respBody utils.Response
	err = json.Unmarshal([]byte(rr.Body.String()), &respBody)
	assert.NoError(t, err)
	assert.Equal(t, respBody.Success, false)
	assert.NotNil(t, respBody.Error)
}
