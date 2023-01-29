package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/internal/create"
)

func RegisterRouter(e *echo.Echo, cs create.Service) *echo.Echo {

	e.GET("/", CreateTask(cs))

	return e
}
