package handler

import (
	"net/http"

	"github.com/Dosugamea/go-todo-api-otel/internal/usecase"
	"github.com/labstack/echo/v4"
)

type TaskHandler interface {
	Add() echo.HandlerFunc
	Get() echo.HandlerFunc
	List() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type taskHandler struct {
	uc usecase.TaskUsecase
}

func NewTaskHandler(uc usecase.TaskUsecase) TaskHandler {
	return taskHandler{
		uc: uc,
	}
}

// CreateTask godoc
// @Summary Create a task
// @Description Create a task.
// @ID create-task
// @Tags task
// @Accept  json
// @Produce  json
// @Param task body response.AddTaskRequest true "Task to create"
// @Success 201 {object} response.AddTaskResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 422 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /tasks [post]
func (h taskHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

// GetTask godoc
// @Summary Get a task
// @Description Get a task.
// @ID get-task
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the task to get"
// @Success 200 {object} response.GetTaskResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /tasks/{id} [get]
func (h taskHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	}
}

// ListTasks godoc
// @Summary Get all tasks
// @Description Get all tasks.
// @ID get-tasks
// @Tags task
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ListTaskResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /tasks [get]
func (h taskHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

// UpdateTask godoc
// @Summary Update a task
// @Description Update a task
// @ID update-task
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the task to update"
// @Param task body response.UpdateTaskRequest true "Task to update"
// @Success 200 {object} response.UpdateTaskResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 422 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /tasks/{id} [put]
func (h taskHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task
// @ID delete-task
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the task to delete"
// @Success 204
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /tasks/{id} [delete]
func (h taskHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
