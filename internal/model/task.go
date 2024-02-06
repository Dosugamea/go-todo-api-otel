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

func NewTask(title string, description string) *Task {
	currentTime := time.Now()
	return &Task{
		Title:       title,
		Description: description,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
		IsCompleted: false,
	}
}
