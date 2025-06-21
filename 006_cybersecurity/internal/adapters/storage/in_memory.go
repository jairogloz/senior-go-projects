package storage

import (
	"errors"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/006_cybersecurity/internal/domain"
	"sync"
)

// InMemoryUserStorage is an in-memory implementation of UserStorage
type InMemoryUserStorage struct {
	mu     sync.Mutex
	users  map[int]domain.User
	nextID int
}

// NewInMemoryUserStorage creates a new instance of InMemoryUserStorage
func NewInMemoryUserStorage() *InMemoryUserStorage {
	return &InMemoryUserStorage{
		users:  make(map[int]domain.User),
		nextID: 1,
	}
}

// Create adds a new user to the storage
func (s *InMemoryUserStorage) Create(user domain.User) (domain.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = s.nextID
	s.nextID++
	s.users[user.ID] = user
	return user, nil
}

// GetByID retrieves a user by ID
func (s *InMemoryUserStorage) GetByID(id int) (domain.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}
