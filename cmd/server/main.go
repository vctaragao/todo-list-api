package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/api/http"
	"github.com/vctaragao/todo-list-api/internal"
	"github.com/vctaragao/todo-list-api/storage"
)

func main() {
	repo := storage.NewPostgress()

	tl := internal.NewTodoList(repo)

	e := http.RegisterRouter(echo.New(), tl)

	e.Logger.Fatal(e.Start(":1323"))
}
