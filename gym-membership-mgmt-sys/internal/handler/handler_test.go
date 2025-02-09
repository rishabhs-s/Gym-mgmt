package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	storage "gym-membership-mgmt-sys/internal/data"
	"gym-membership-mgmt-sys/internal/models"
)


func TestRegisterMembership(t *testing.T) {

	reqBody := `{"name": "Joe", "email": "rs@gm.com", "start_date": "2025-02-10"}`

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	RegisterMembership(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestViewMembership(t *testing.T) {
	storage.Store.RegisterMembership(models.Membership{
		Name:      "Joe",
		Email:     "rs@gm.com",
		StartDate: models.Date{Time: time.Date(2025, 02, 10, 0, 0, 0, 0, time.UTC)},
	})

	req := httptest.NewRequest("GET", "/membership?email=rs@gm.com", nil)
	rr := httptest.NewRecorder()

	ViewMembership(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestViewAllActiveMembers(t *testing.T) {
	storage.Store.RegisterMembership(models.Membership{Name: "RS", Email: "rs@gm.com"})
	storage.Store.RegisterMembership(models.Membership{Name: "RS1", Email: "r1s@gm.com"})

	req := httptest.NewRequest("GET", "/members", nil)
	rr := httptest.NewRecorder()

	ViewAllActiveMembers(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestModifyStartDate(t *testing.T) {
	storage.Store.RegisterMembership(models.Membership{Name: "RS", Email: "rs@gm.com"})

	reqBody := `{"email": "rs@gm.com", "new_date": "2025-03-01"}`
	req := httptest.NewRequest("PUT", "/modify", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	ModifyStartDate(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestModifyStartDate_InvalidDateFormat(t *testing.T) {
	storage.Store.RegisterMembership(models.Membership{Name: "RS", Email: "rs@gm.com"})

	reqBody := `{"email": "rs@gm.com", "new_date": "xxx-01-02"}`
	req := httptest.NewRequest("PUT", "/modify", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	ModifyStartDate(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}