package response

import (
	"errors"
	"time"

	"github.com/Dosugamea/go-todo-api-otel/internal/model"
)

type TaskResponse struct {
	ID          int    `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func (t *TaskResponse) Bind(task *model.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}
	t.ID = task.ID
	t.CreatedAt = task.CreatedAt.Format(time.RFC3339)
	t.UpdatedAt = task.UpdatedAt.Format(time.RFC3339)
	t.Title = task.Title
	t.Description = task.Description
	t.IsCompleted = task.IsCompleted
	return nil
}

type ListTaskResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}

func (r *ListTaskResponse) Bind(tasks []*model.Task) error {
	newTasks := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		res := TaskResponse{}
		if err := res.Bind(task); err != nil {
			return err
		}
		newTasks[i] = res
	}
	r.Tasks = newTasks
	return nil
}

type GetTaskResponse struct {
	Task TaskResponse `json:"task"`
}

func (r *GetTaskResponse) Bind(task *model.Task) error {
	res := TaskResponse{}
	if err := res.Bind(task); err != nil {
		return err
	}
	r.Task = res
	return nil
}

type CreateTaskResponse struct {
	Task TaskResponse `json:"task"`
}

func (r *CreateTaskResponse) Bind(task *model.Task) error {
	res := TaskResponse{}
	if err := res.Bind(task); err != nil {
		return err
	}
	r.Task = res
	return nil
}

type UpdateTaskResponse struct {
	Task TaskResponse `json:"task"`
}

func (r *UpdateTaskResponse) Bind(task *model.Task) error {
	res := TaskResponse{}
	if err := res.Bind(task); err != nil {
		return err
	}
	r.Task = res
	return nil
}
