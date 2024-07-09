package model

import (
	"database/sql"
	"time"
)

// User represents a user in the business logic layer.
type User struct {
	ID        int64
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// CreateUser represents the data needed to create a new user.
type CreateUser struct {
	Name            string
	Email           string
	Password        string
	Role            string
	ConfirmPassword string
}

// UpdateUser represents the data needed to update an existing user.
type UpdateUser struct {
	ID    int64
	Name  string
	Email string
	Role  string
}
