package usecase

import (
	"github.com/Dosugamea/go-todo-api-otel/internal/model"
	"github.com/Dosugamea/go-todo-api-otel/internal/repository"
)

type TaskUsecase interface {
	Get(id int) (*model.Task, error)
	List() ([]*model.Task, error)
	Create(task *model.Task) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return taskUsecase{
		repo: repo,
	}
}

func (uc taskUsecase) Get(id int) (*model.Task, error) {
	return uc.repo.FindByID(id)
}

func (uc taskUsecase) List() ([]*model.Task, error) {
	return uc.repo.FindAll()
}

func (uc taskUsecase) Create(task *model.Task) (*model.Task, error) {
	resp, err := uc.repo.Create(task)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc taskUsecase) Update(task *model.Task) (*model.Task, error) {
	if err := uc.repo.Update(task); err != nil {
		return nil, err
	}
	resp, err := uc.repo.FindByID(task.ID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc taskUsecase) Delete(id int) error {
	if err := uc.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
