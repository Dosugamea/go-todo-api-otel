package repository

import (
	"context"

	"github.com/Dosugamea/go-todo-api-otel/internal/model"
)

type TaskRepository interface {
	FindAll(ctx context.Context) ([]*model.Task, error)
	FindByID(ctx context.Context, id int) (*model.Task, error)
	Create(ctx context.Context, task *model.Task) (*model.Task, error)
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id int) error
}
