package handler

import (
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

func (h taskHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h taskHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h taskHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h taskHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (h taskHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
