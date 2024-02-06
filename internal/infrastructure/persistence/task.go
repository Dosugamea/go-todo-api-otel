package persistence

import (
	"context"
	"time"

	"github.com/Dosugamea/go-todo-api-otel/internal/infrastructure/observability"
	"github.com/Dosugamea/go-todo-api-otel/internal/model"
	"github.com/Dosugamea/go-todo-api-otel/internal/repository"
	"gorm.io/gorm"
)

type taskRepositoryImpl struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return taskRepositoryImpl{
		Conn: conn,
	}
}

func (r taskRepositoryImpl) FindAll(ctx context.Context) ([]*model.Task, error) {
	ctx, span := observability.Tracer.StartPersistenceSpan(ctx, "FindAll")
	defer span.End()

	var tasks []*model.Task
	result := r.Conn.WithContext(ctx).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r taskRepositoryImpl) FindByID(ctx context.Context, id int) (*model.Task, error) {
	ctx, span := observability.Tracer.StartPersistenceSpan(ctx, "FindByID")
	defer span.End()

	// 1~3秒のランダムな待ち時間を追加
	time.Sleep(time.Duration(1+time.Now().Nanosecond()%3) * time.Second)

	var task model.Task
	if result := r.Conn.WithContext(ctx).Model(&model.Task{}).Where("id = ?", id).Take(&task); result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r taskRepositoryImpl) Create(ctx context.Context, task *model.Task) (*model.Task, error) {
	ctx, span := observability.Tracer.StartPersistenceSpan(ctx, "Create")
	defer span.End()

	if err := r.Conn.WithContext(ctx).Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r taskRepositoryImpl) Update(ctx context.Context, task *model.Task) error {
	ctx, span := observability.Tracer.StartPersistenceSpan(ctx, "Update")
	defer span.End()

	time.Sleep(3 * time.Second)

	if err := r.Conn.WithContext(ctx).Model(&task).Updates(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r taskRepositoryImpl) Delete(ctx context.Context, id int) error {
	ctx, span := observability.Tracer.StartPersistenceSpan(ctx, "Delete")
	defer span.End()

	if err := r.Conn.WithContext(ctx).Model(&model.Task{}).Where("id = ?", id).Delete(&model.Task{}).Error; err != nil {
		return err
	}
	return nil
}
