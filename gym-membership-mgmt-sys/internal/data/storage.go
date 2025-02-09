package storage

import (
    "sync"
    "gym-membership-mgmt-sys/internal/models"
)
type MembershipStore struct {
    mu   sync.Mutex
    data map[string]models.Membership 
}

var Store = &MembershipStore{
    data: make(map[string]models.Membership),
}

func (s *MembershipStore) RegisterMembership(m models.Membership) models.Membership {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[m.Email] = m
    return m
}

func (s *MembershipStore) GetMembership(email string) (models.Membership, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    m, exists := s.data[email]
    return m, exists
}

func (s *MembershipStore) GetAllActiveMembers() []models.Membership {
    s.mu.Lock()
    defer s.mu.Unlock()
    members := []models.Membership{}
    for _, m := range s.data {
        members = append(members, m)
    }
    return members
}

func (s *MembershipStore) CancelMembership(email string) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    if _, exists := s.data[email]; exists {
        delete(s.data, email)
        return true
    }
    return false
}

func (s *MembershipStore) ModifyStartDate(email string, newDate models.Date) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    if m, exists := s.data[email]; exists {
        m.StartDate = newDate
        s.data[email] = m
        return true
    }
    return false
}
