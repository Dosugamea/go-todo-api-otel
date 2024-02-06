package model

import "time"

type Task struct {
	ID          int       `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Title       string
	Description string
	IsCompleted bool
}

func (t *Task) ChangeData(title string, description string, isCompleted bool) {
	t.Title = title
	t.Description = description
	t.IsCompleted = isCompleted
	t.UpdatedAt = time.Now()
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
