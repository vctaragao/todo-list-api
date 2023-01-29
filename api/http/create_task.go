package http

import (
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/internal/create"
)

type TaskDto struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

func CreateTask(cs create.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		dto := new(TaskDto)

		if err := ctx.Bind(dto); err != nil {
			return err
		}

		t := create.NewTaskDto(dto.Description, dto.Priority)

		id, _ := cs.Create(t)

		return ctx.String(http.StatusOK, strconv.Itoa(id))
	}
}
