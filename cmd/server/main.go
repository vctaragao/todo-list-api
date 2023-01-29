package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/api/http"
	"github.com/vctaragao/todo-list-api/internal/create"
	"github.com/vctaragao/todo-list-api/storage"
)

func main() {
	repo := storage.NewDummyAdapter()

	cs := create.CreateService(repo)

	e := http.RegisterRouter(echo.New(), cs)

	e.Logger.Fatal(e.Start(":1323"))
}
