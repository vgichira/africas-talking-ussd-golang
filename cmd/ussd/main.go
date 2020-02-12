package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	port := "8080"

	router.HandleFunc("/ussd", ussd).Methods("POST")

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server running on port " + port)

	log.Fatal(srv.ListenAndServe())
}
