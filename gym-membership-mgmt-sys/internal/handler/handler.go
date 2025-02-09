package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	storage "gym-membership-mgmt-sys/internal/data"
	"gym-membership-mgmt-sys/internal/models"
)

func RegisterMembership(w http.ResponseWriter, r *http.Request) {
	var m models.Membership
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Print("Invalid input", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	registeredMember := storage.Store.RegisterMembership(m)
	json.NewEncoder(w).Encode(registeredMember)
	log.Print(registeredMember)
}

func ViewMembership(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if m, exists := storage.Store.GetMembership(email); exists {
		json.NewEncoder(w).Encode(m)
		log.Print("MemberShip Details", m)
	} else {
		http.Error(w, "Membership not found", http.StatusNotFound)
		log.Print("Membership not found")
	}
}

func ViewAllActiveMembers(w http.ResponseWriter, r *http.Request) {
	members := storage.Store.GetAllActiveMembers()
	if len(members) == 0 {
		log.Print("No Reg members:")
		json.NewEncoder(w).Encode(members)
	} else {
		json.NewEncoder(w).Encode(members)
		log.Print("Reg members: ", members)
	}

}

func CancelMembership(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if storage.Store.CancelMembership(email) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Membership cancelled successfully for : %s", email)))
		log.Print("Member ship cancelled for ", email)

	} else {
		log.Printf("Member doesnt exist")
		http.Error(w, "Membership not found", http.StatusNotFound)
	}
}

func ModifyStartDate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email   string `json:"email"`
		NewDate string `json:"new_date"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	parsedDate, err := time.Parse("2006-01-02", req.NewDate)
	if err != nil {
		http.Error(w, "Invalid date format, expected YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	log.Printf("Modifying start date for %s to %s", req.Email, req.NewDate)

	if storage.Store.ModifyStartDate(req.Email, models.Date{Time: parsedDate}) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Membership start date updated successfully"))
		log.Printf("Modified start date for %s to %s", req.Email, req.NewDate)
	} else {
		log.Printf("Email doesnt exist")
		http.Error(w, "Membership not found", http.StatusNotFound)
	}
}
