package transport

import (
	"encoding/json"
	"net/http"

	"github.com/chidam1994/happyfox/group"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type groupHandler struct {
	groupSvc group.Service
}

func (handler *groupHandler) createGroup(w http.ResponseWriter, r *http.Request) {
	var body *CreateGroupRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w = utils.GetBadReqResponse(w, err.Error())
		return
	}
	group, err := body.Validate()
	if err != nil {
		w = utils.GetBadReqResponse(w, err.Error())
		return
	}
	id, err := handler.groupSvc.SaveGroup(group)
	if err != nil {
		w = utils.GetFailureResponse(w, err)
		return
	}
	data := &GroupId{Id: id.String()}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		w = utils.GetBadReqResponse(w, err.Error())
		w.Write([]byte(err.Error()))
	}
}

func (handler *groupHandler) deleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr := vars["id"]
	id, err := uuid.Parse(idstr)
	if err != nil {
		w = utils.GetBadReqResponse(w, "invalid uuid in request")
		return
	}
	err = handler.groupSvc.DeleteGroup(id)
	if err != nil {
		w = utils.GetFailureResponse(w, err)
		return
	}
	w = utils.GetSuccessReqResponse(w)
}

func (handler *groupHandler) getGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr := vars["id"]
	id, err := uuid.Parse(idstr)
	if err != nil {
		w = utils.GetBadReqResponse(w, "invalid uuid in request")
		return
	}
	group, err := handler.groupSvc.GetGroup(id)
	if err != nil {
		w = utils.GetFailureResponse(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(group); err != nil {
		w = utils.GetBadReqResponse(w, err.Error())
		w.Write([]byte(err.Error()))
	}
}

func (handler *groupHandler) updateGroup(w http.ResponseWriter, r *http.Request) {
	var body UpdateGroupRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w = utils.GetBadReqResponse(w, err.Error())
		return
	}
	vars := mux.Vars(r)
	actionStr := vars["action"]
	action, err := getUpdateAction(actionStr)
	if err != nil {
		w = utils.GetBadReqResponse(w, err.Error())
		return
	}
	groupId, err := uuid.Parse(body.GroupId)
	if err != nil {
		w = utils.GetBadReqResponse(w, "invalid groupId in request")
		return
	}

	switch action {
	case addMembers:
		memberIds, err := validateUuids(body.MemberIds)
		if err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			return
		}
		err = handler.groupSvc.AddMembers(groupId, memberIds)
		if err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			return
		}
	case remMembers:
		memberIds, err := validateUuids(body.MemberIds)
		if err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			return
		}
		err = handler.groupSvc.RemMembers(groupId, memberIds)
		if err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			return
		}
	case renameGroup:
		if body.Name == "" {
			w = utils.GetBadReqResponse(w, "name cannot be empty in request")
			return
		}
		err = handler.groupSvc.RenameGroup(groupId, body.Name)
		if err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			return
		}
	}
	w = utils.GetSuccessReqResponse(w)
}

func InitGroupHandlers(r *mux.Router, service group.Service) {
	handler := &groupHandler{
		groupSvc: service,
	}
	r.HandleFunc("", handler.createGroup).Methods("POST", "OPTIONS")
	r.HandleFunc("/{id}", handler.deleteGroup).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/{id}", handler.getGroup).Methods("GET", "OPTIONS")
	r.HandleFunc("/{action}", handler.updateGroup).Methods("PUT", "OPTIONS")
}
