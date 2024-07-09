package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type CreateUser struct {
	Name            string
	Email           string
	Password        string
	Role            string
	ConfirmPassword string
}

type UpdateUser struct {
	ID    int64
	Name  string
	Email string
	Role  string
}
