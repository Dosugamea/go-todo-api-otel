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
	var tasks []*model.Task
	result := r.Conn.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r taskRepositoryImpl) FindByID(id int) (*model.Task, error) {
	var task model.Task
	if result := r.Conn.Model(&model.Task{}).Where("id = ?", id).Take(&task); result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r taskRepositoryImpl) Create(task *model.Task) (*model.Task, error) {
	if err := r.Conn.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r taskRepositoryImpl) Update(task *model.Task) error {
	if err := r.Conn.Model(&task).Updates(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r taskRepositoryImpl) Delete(id int) error {
	if err := r.Conn.Model(&model.Task{}).Where("id = ?", id).Delete(&model.Task{}).Error; err != nil {
		return err
	}
	return nil
}
