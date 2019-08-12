package transport

import (
	"encoding/json"
	"net/http"

	"github.com/chidam1994/happyfox/contact"
	"github.com/chidam1994/happyfox/models"
	"github.com/chidam1994/happyfox/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type contactHandler struct {
	contactSvc contact.Service
}

func (handler *contactHandler) createContact(w http.ResponseWriter, r *http.Request) {
	var body *CreateContactRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w = utils.GetBadReqResponse(w, err.Error())
		return
	}
	contact, err := body.Validate()
	if err != nil {
		w = utils.GetBadReqResponse(w, err.Error())
		return
	}
	id, err := handler.contactSvc.SaveContact(contact)
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

func (handler *contactHandler) deleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr := vars["id"]
	id, err := uuid.Parse(idstr)
	if err != nil {
		w = utils.GetBadReqResponse(w, "invalid uuid in request")
		return
	}
	err = handler.contactSvc.DeleteContact(id)
	if err != nil {
		w = utils.GetFailureResponse(w, err)
		return
	}
	w = utils.GetSuccessReqResponse(w)
}

func (handler *contactHandler) getContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr := vars["id"]
	id, err := uuid.Parse(idstr)
	if err != nil {
		w = utils.GetBadReqResponse(w, "invalid uuid in request")
		return
	}
	contact, err := handler.contactSvc.GetContact(id)
	if err != nil {
		w = utils.GetFailureResponse(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(contact); err != nil {
		w = utils.GetBadReqResponse(w, err.Error())
		w.Write([]byte(err.Error()))
	}
}

func (handler *contactHandler) searchContact(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	filtersMap := make(map[models.Filter]string)
	for key, val := range values {
		filter, err := models.GetFilter(key)
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
	results, err := handler.contactSvc.FindContacts(filtersMap)
	if err != nil {
		w = utils.GetFailureResponse(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(results); err != nil {
		w = utils.GetBadReqResponse(w, err.Error())
		w.Write([]byte(err.Error()))
	}
}

func InitContactHandlers(r *mux.Router, service contact.Service) {
	handler := &contactHandler{
		contactSvc: service,
	}
	r.HandleFunc("", handler.createContact).Methods("POST", "OPTIONS")
	r.HandleFunc("/{id}", handler.deleteContact).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/search", handler.searchContact).Methods("GET", "OPTIONS")
	r.HandleFunc("/{id}", handler.getContact).Methods("GET", "OPTIONS")
}
