package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/internal"
)

func RegisterRouter(e *echo.Echo, tl internal.TodoList) *echo.Echo {

	e.GET("/", CreateTask(tl))

	return e
}
