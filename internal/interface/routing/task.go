package routing

import (
	"github.com/Dosugamea/go-todo-api-otel/internal/interface/handler"
	"github.com/labstack/echo/v4"
)

func RegisterTaskRoutings(g *echo.Group, h handler.TaskHandler) {
	g.POST("/tasks", h.Create)
	g.GET("/tasks/:id", h.Get)
	g.GET("/tasks", h.List)
	g.PUT("/tasks/:id", h.Update)
	g.DELETE("/tasks/:id", h.Delete)
}
