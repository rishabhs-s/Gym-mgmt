package main

import (
	"log"
	"net/http"

	"gym-membership-mgmt-sys/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Entering Main")
	r := mux.NewRouter()

	r.HandleFunc("/register", handler.RegisterMembership).Methods("POST")
	r.HandleFunc("/membership", handler.ViewMembership).Methods("GET")
	r.HandleFunc("/members", handler.ViewAllActiveMembers).Methods("GET")
	r.HandleFunc("/cancel", handler.CancelMembership).Methods("DELETE")
	r.HandleFunc("/modify", handler.ModifyStartDate).Methods("PUT")
    log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
