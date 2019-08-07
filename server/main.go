package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chidam1994/happyfox/config"
	"github.com/chidam1994/happyfox/datastore/gorp"
	"github.com/chidam1994/happyfox/server/handlers"
	"github.com/chidam1994/happyfox/services/contactsvc"
	"github.com/chidam1994/happyfox/services/groupsvc"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	dbMap := gorp.InitDB()
	contactService := contactsvc.NewService(contactsvc.NewPgsqlRepo(dbMap))
	groupService := groupsvc.NewService(groupsvc.NewPgsqlRepo(dbMap))
	defer gorp.CloseDBConn()
	contactRouter := r.PathPrefix("/contact").Subrouter()
	groupRouter := r.PathPrefix("/group").Subrouter()
	handlers.InitContactHandlers(contactRouter, contactService)
	handlers.InitGroupHandlers(groupRouter, groupService)

	http.Handle("/", r)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         config.GetString(config.PORT),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
