package types

import (
	"time"
)

type Todo struct {
	ID          int
	CreatedAt   time.Time
	Description string
	UpdatedAt   time.Time
	UserID      int
}

type User struct {
	ID        int
	CreatedAt time.Time
	Email     string
	Name      string
	Password  string
	UpdatedAt time.Time
}
