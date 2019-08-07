package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chidam1994/happyfox/services/contactsvc"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func createContact(svc *contactsvc.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var body *CreateContactRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			w = utils.GetBadReqResponse(w, "error decoding json body")
			return
		}
		contact, err := body.Validate()
		if err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			return
		}
		id, err := svc.SaveContact(contact)
		if err != nil {
			w = utils.GetFailureResponse(w, err)
			return
		}
		data := &ContactId{Id: id.String()}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			w.Write([]byte(err.Error()))
		}
	}
}

func deleteContact(svc *contactsvc.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idstr := vars["id"]
		id, err := uuid.Parse(idstr)
		if err != nil {
			w = utils.GetBadReqResponse(w, "invalid uuid in request")
			return
		}
		err = svc.DeleteContact(id)
		if err != nil {
			w = utils.GetFailureResponse(w, err)
			return
		}
		w = utils.GetSuccessReqResponse(w)
	}
}

func getContact(svc *contactsvc.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idstr := vars["id"]
		id, err := uuid.Parse(idstr)
		if err != nil {
			w = utils.GetBadReqResponse(w, "invalid uuid in request")
			return
		}
		contact, err := svc.GetContact(id)
		if err != nil {
			w = utils.GetFailureResponse(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(contact); err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			w.Write([]byte(err.Error()))
		}
	}
}

func searchContact(svc *contactsvc.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		filtersMap := make(map[contactsvc.Filter]string)
		for key, val := range values {
			filter, err := contactsvc.GetFilter(key)
			if err != nil {
				w = utils.GetBadReqResponse(w, "unrecognized query params present in request")
				return
			}
			if len(val) > 1 {
				w = utils.GetBadReqResponse(w, "only one value can be specified for a filter")
				return
			}
			filtersMap[filter] = val[0]
		}
		results, err := svc.FindContacts(filtersMap)
		if err != nil {
			w = utils.GetFailureResponse(w, err)
			return
		}
		if err := json.NewEncoder(w).Encode(results); err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			w.Write([]byte(err.Error()))
		}
	}
}
func InitContactHandlers(r *mux.Router, service *contactsvc.Service) {
	r.HandleFunc("", createContact(service)).Methods("POST", "OPTIONS")
	r.HandleFunc("/{id}", deleteContact(service)).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/search", searchContact(service)).Methods("GET", "OPTIONS")
	r.HandleFunc("/{id}", getContact(service)).Methods("GET", "OPTIONS")
}
