package domain

import "time"

// User represents a user in the system
type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Age       int        `json:"age"`
	CreatedAt *time.Time `json:"createdAt"`
	IsAdmin   bool       `json:"isAdmin"`
}

// UserStorage defines the interface for user storage
type UserStorage interface {
	Create(user User) (User, error)
	GetByID(id int) (User, error)
}

// UserService defines the interface for user-related operations
type UserService interface {
	Create(user User) (User, error)
	GetByID(id int) (User, error)
}
