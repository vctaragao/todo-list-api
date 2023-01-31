package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/internal/create_task"
)

func RegisterRouter(e *echo.Echo, ct create_task.Service) *echo.Echo {

	e.GET("/", CreateTask(ct))

	return e
}
