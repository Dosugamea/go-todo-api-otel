package model

import "time"

type Task struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description string
	IsCompleted bool
}
