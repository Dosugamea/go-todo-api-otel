package repository

import "github.com/Dosugamea/go-todo-api-otel/internal/model"

type TaskRepository interface {
	FindAll() ([]*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Create(task *model.Task) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
}
