package oauth

import (
	"sync"
	"time"
)

type stateEntry struct {
	value     string
	expiresAt time.Time
}

// StateStore manages OAuth state tokens with TTL
type StateStore struct {
	mu     sync.RWMutex
	states map[string]stateEntry
}

// NewStateStore creates a new state store
func NewStateStore() *StateStore {
	store := &StateStore{
		states: make(map[string]stateEntry),
	}
	// Start cleanup goroutine
	go store.cleanup()
	return store
}

// Set stores a state with TTL
func (s *StateStore) Set(state string, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.states[state] = stateEntry{
		value:     state,
		expiresAt: time.Now().Add(ttl),
	}
}

// Validate checks if state exists and is valid, then removes it
func (s *StateStore) Validate(state string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, exists := s.states[state]
	if !exists {
		return false
	}

	// Check if expired
	if time.Now().After(entry.expiresAt) {
		delete(s.states, state)
		return false
	}

	// Valid state, remove it (one-time use)
	delete(s.states, state)
	return true
}

// cleanup removes expired states periodically
func (s *StateStore) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		s.mu.Lock()
		now := time.Now()
		for state, entry := range s.states {
			if now.After(entry.expiresAt) {
				delete(s.states, state)
			}
		}
		s.mu.Unlock()
	}
}
