package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chidam1994/happyfox/services/contactsvc"
	"github.com/chidam1994/happyfox/utils"
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
			w = utils.GetBadReqResponse(w, err.Error())
			return
		}
		data := &CreateContactResponse{Id: id.String()}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w = utils.GetBadReqResponse(w, err.Error())
			w.Write([]byte(err.Error()))
		}
	}
}

func InitContactHandlers(r *mux.Router, service *contactsvc.Service) {
	r.HandleFunc("/create", createContact(service)).Methods("POST", "OPTIONS")
}
