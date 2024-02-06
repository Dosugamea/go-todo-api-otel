package request

import (
	"github.com/labstack/echo/v4"
)

type AddTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *AddTaskRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	// TODO: バリデーション
	return nil
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func (r *UpdateTaskRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	// TODO: バリデーション
	return nil
}
