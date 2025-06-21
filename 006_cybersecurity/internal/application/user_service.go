package application

import (
	"errors"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/006_cybersecurity/internal/domain"
	"time"
)

// UserServiceImpl is an implementation of UserService
type UserServiceImpl struct {
	storage domain.UserStorage
}

// NewUserService creates a new instance of UserServiceImpl
func NewUserService(storage domain.UserStorage) *UserServiceImpl {
	return &UserServiceImpl{storage: storage}
}

// Create creates a new user with validation
func (s *UserServiceImpl) Create(user domain.User) (domain.User, error) {
	if user.Name == "" {
		return domain.User{}, errors.New("name cannot be empty")
	}
	if user.Age < 18 {
		return domain.User{}, errors.New("age must be greater than 18")
	}

	now := time.Now()
	user.CreatedAt = &now

	return s.storage.Create(user)
}

// GetByID retrieves a user by ID
func (s *UserServiceImpl) GetByID(id int) (domain.User, error) {
	return s.storage.GetByID(id)
}
