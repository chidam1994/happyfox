package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chidam1994/happyfox/server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	contactRouter := r.PathPrefix("/contact").Subrouter()
	handlers.InitContactHandlers(contactRouter)

	http.Handle("/", r)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
