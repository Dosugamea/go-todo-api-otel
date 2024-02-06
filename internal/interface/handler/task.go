package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Dosugamea/go-todo-api-otel/internal/interface/handler/request"
	"github.com/Dosugamea/go-todo-api-otel/internal/interface/handler/response"
	"github.com/Dosugamea/go-todo-api-otel/internal/model"
	"github.com/Dosugamea/go-todo-api-otel/internal/usecase"
	shared_err "github.com/Dosugamea/go-todo-api-otel/pkg/shared/err"
	"github.com/labstack/echo/v4"
)

type TaskHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	List(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
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
// @Param task body request.CreateTaskRequest true "Task to create"
// @Success 201 {object} response.CreateTaskResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 422 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /tasks [post]
func (h taskHandler) Create(c echo.Context) error {
	req := &request.CreateTaskRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err))
	}

	task := model.NewTask(req.Title, req.Description)
	createdTask, err := h.uc.Create(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(shared_err.ErrInternalServerError))
	}

	res := response.CreateTaskResponse{}
	if err := res.Bind(createdTask); err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err))
	}
	return c.JSON(http.StatusCreated, res)
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
func (h taskHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(shared_err.ErrInvalidId))
	}

	task, err := h.uc.Get(id)
	if err != nil {
		if v := errors.Is(err, shared_err.ErrNotFound); v {
			return c.JSON(http.StatusNotFound, response.NewErrorResponse(shared_err.ErrNotFound))
		}
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(shared_err.ErrInternalServerError))
	}

	res := response.GetTaskResponse{}
	if err := res.Bind(task); err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err))
	}
	return c.JSON(http.StatusOK, res)
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
func (h taskHandler) List(c echo.Context) error {
	tasks, err := h.uc.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(shared_err.ErrInternalServerError))
	}

	res := response.ListTaskResponse{}
	if err := res.Bind(tasks); err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(err))
	}
	return c.JSON(http.StatusOK, res)
}

// UpdateTask godoc
// @Summary Update a task
// @Description Update a task
// @ID update-task
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the task to update"
// @Param task body request.UpdateTaskRequest true "Task to update"
// @Success 200 {object} response.UpdateTaskResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 422 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /tasks/{id} [put]
func (h taskHandler) Update(c echo.Context) error {
	req := &request.UpdateTaskRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.NewErrorResponse(err))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(shared_err.ErrInvalidId))
	}

	updatedTask, err := h.uc.Update(id, req.Title, req.Description, req.IsCompleted)
	if err != nil {
		if v := errors.Is(err, shared_err.ErrNotFound); v {
			return c.JSON(http.StatusNotFound, response.NewErrorResponse(shared_err.ErrNotFound))
		}
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(shared_err.ErrInternalServerError))
	}

	res := response.UpdateTaskResponse{}
	if err := res.Bind(updatedTask); err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(shared_err.ErrInternalServerError))
	}

	return c.JSON(http.StatusOK, res)
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
func (h taskHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewErrorResponse(shared_err.ErrInvalidId))
	}

	if err := h.uc.Delete(id); err != nil {
		if v := errors.Is(err, shared_err.ErrNotFound); v {
			return c.JSON(http.StatusNotFound, response.NewErrorResponse(shared_err.ErrNotFound))
		}
		return c.JSON(http.StatusInternalServerError, response.NewErrorResponse(shared_err.ErrInternalServerError))
	}

	return c.JSON(http.StatusNoContent, nil)
}
