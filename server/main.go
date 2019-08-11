package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chidam1994/happyfox/config"
	_contactRepo "github.com/chidam1994/happyfox/contact/repository"
	_contactService "github.com/chidam1994/happyfox/contact/service"
	_contactTransport "github.com/chidam1994/happyfox/contact/transport"
	"github.com/chidam1994/happyfox/datastore/gorp"
	_groupRepo "github.com/chidam1994/happyfox/group/repository"
	_groupService "github.com/chidam1994/happyfox/group/service"
	_groupTransport "github.com/chidam1994/happyfox/group/transport"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	dbMap := gorp.InitDB()
	contactService := _contactService.NewService(_contactRepo.NewPgsqlRepo(dbMap))
	groupService := _groupService.NewService(_groupRepo.NewPgsqlRepo(dbMap))
	defer gorp.CloseDBConn()
	contactRouter := r.PathPrefix("/contact").Subrouter()
	groupRouter := r.PathPrefix("/group").Subrouter()
	_contactTransport.InitContactHandlers(contactRouter, contactService)
	_groupTransport.InitGroupHandlers(groupRouter, groupService)

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
