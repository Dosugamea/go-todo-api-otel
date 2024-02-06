package usecase

import (
	"context"

	"github.com/Dosugamea/go-todo-api-otel/internal/infrastructure/observability"
	"github.com/Dosugamea/go-todo-api-otel/internal/model"
	"github.com/Dosugamea/go-todo-api-otel/internal/repository"
)

type TaskUsecase interface {
	Get(ctx context.Context, id int) (*model.Task, error)
	List(ctx context.Context) ([]*model.Task, error)
	Create(ctx context.Context, task *model.Task) (*model.Task, error)
	Update(ctx context.Context, id int, name string, description string, isCompleted bool) (*model.Task, error)
	Delete(ctx context.Context, id int) error
}

type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo repository.TaskRepository) TaskUsecase {
	return taskUsecase{
		repo: repo,
	}
}

func (uc taskUsecase) Get(ctx context.Context, id int) (*model.Task, error) {
	ctx, span := observability.Tracer.StartUsecaseSpan(ctx, "Get")
	defer span.End()

	return uc.repo.FindByID(ctx, id)
}

func (uc taskUsecase) List(ctx context.Context) ([]*model.Task, error) {
	ctx, span := observability.Tracer.StartUsecaseSpan(ctx, "List")
	defer span.End()
	return uc.repo.FindAll(ctx)
}

func (uc taskUsecase) Create(ctx context.Context, task *model.Task) (*model.Task, error) {
	ctx, span := observability.Tracer.StartUsecaseSpan(ctx, "Create")
	defer span.End()
	resp, err := uc.repo.Create(ctx, task)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc taskUsecase) Update(ctx context.Context, id int, name string, description string, isCompleted bool) (*model.Task, error) {
	ctx, span := observability.Tracer.StartUsecaseSpan(ctx, "Update")
	defer span.End()
	task, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	task.ChangeData(name, description, isCompleted)

	if err := uc.repo.Update(ctx, task); err != nil {
		return nil, err
	}

	resp, err := uc.repo.FindByID(ctx, task.ID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc taskUsecase) Delete(ctx context.Context, id int) error {
	ctx, span := observability.Tracer.StartUsecaseSpan(ctx, "Delete")
	defer span.End()
	if err := uc.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
