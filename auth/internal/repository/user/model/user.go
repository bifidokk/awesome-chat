package model

import (
	"database/sql"
	"time"
)

// User represents a user record in the data store.
type User struct {
	ID        int64
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
