package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chidam1994/happyfox/services/groupsvc"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func createGroup(svc *groupsvc.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body *CreateGroupRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			w = utils.GetBadReqResponse(w, "error decoding json body")
			return
		}
		group, err := body.Validate()
		if err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			return
		}
		id, err := svc.SaveGroup(group)
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
}

func deleteGroup(svc *groupsvc.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idstr := vars["id"]
		id, err := uuid.Parse(idstr)
		if err != nil {
			w = utils.GetBadReqResponse(w, "invalid uuid in request")
			return
		}
		err = svc.DeleteGroup(id)
		if err != nil {
			w = utils.GetFailureResponse(w, err)
			return
		}
		w = utils.GetSuccessReqResponse(w)
	}
}

func InitGroupHandlers(r *mux.Router, service *groupsvc.Service) {
	r.HandleFunc("", createGroup(service)).Methods("POST", "OPTIONS")
	r.HandleFunc("/{id}", deleteGroup(service)).Methods("DELETE", "OPTIONS")
}
