package http

import (
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/internal"
)

type TaskDto struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

func CreateTask(tl internal.TodoList) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		dto := new(TaskDto)

		if err := ctx.Bind(dto); err != nil {
			return err
		}

		id, err := tl.CreateTask(dto.Description, dto.Priority)

		if err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}

		return ctx.String(http.StatusOK, strconv.Itoa(id))
	}
}
