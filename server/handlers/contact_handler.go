package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func contactFind(w http.ResponseWriter, r *http.Request) {
}

func InitContactHandlers(r *mux.Router) {
	r.HandleFunc("/{id}", contactFind).Methods("GET", "OPTIONS")
}
