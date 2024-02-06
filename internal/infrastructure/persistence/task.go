package persistence

import (
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

func (r taskRepositoryImpl) FindAll() ([]*model.Task, error) {
	return nil, nil
}

func (r taskRepositoryImpl) FindByID(id int) (*model.Task, error) {
	return nil, nil
}

func (r taskRepositoryImpl) Create(task *model.Task) (*model.Task, error) {
	return nil, nil
}

func (r taskRepositoryImpl) Update(task *model.Task) error {
	return nil
}

func (r taskRepositoryImpl) Delete(id int) error {
	return nil
}
