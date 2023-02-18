package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/internal"
)

func RegisterRouter(e *echo.Echo, tl internal.TodoList) *echo.Echo {

	e.POST("/create", CreateTask(tl))

	return e
}
