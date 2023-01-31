package http

import (
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	cT "github.com/vctaragao/todo-list-api/internal/create_task"
)

type TaskDto struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

func CreateTask(ct cT.Service) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		dto := new(TaskDto)

		if err := ctx.Bind(dto); err != nil {
			return err
		}

		t := cT.NewTaskDto(dto.Description, dto.Priority)

		id, _ := ct.Create(t)

		return ctx.String(http.StatusOK, strconv.Itoa(id))
	}
}
