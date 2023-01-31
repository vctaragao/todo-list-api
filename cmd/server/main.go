package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/api/http"
	"github.com/vctaragao/todo-list-api/internal/create_task"
	"github.com/vctaragao/todo-list-api/storage"
)

func main() {
	repo := storage.NewDummyAdapter()

	ct := create_task.NewServie(repo)

	e := http.RegisterRouter(echo.New(), ct)

	e.Logger.Fatal(e.Start(":1323"))
}
