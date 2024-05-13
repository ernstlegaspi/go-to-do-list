package types

import (
	"time"
)

type Todo struct {
	ID          int
	CreatedAt   time.Time
	Description string
	UpdatedAt   time.Time
}
