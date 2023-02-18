package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/api/http"
	"github.com/vctaragao/todo-list-api/internal"
	"github.com/vctaragao/todo-list-api/storage"
)

const SQLITE_FILE_PATH = "todo_list.db"

func main() {
	repo := storage.NewSqLite(SQLITE_FILE_PATH)

	ct := internal.NewTodoList(repo)

	e := http.RegisterRouter(echo.New(), ct)

	e.Logger.Fatal(e.Start(":1323"))
}
