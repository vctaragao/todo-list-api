package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/api/http"
	"github.com/vctaragao/todo-list-api/internal"
	"github.com/vctaragao/todo-list-api/storage"
)

func main() {
	dbHost := getDbHost()
	repo := storage.NewSqLite(dbHost)

	tl := internal.NewTodoList(repo)

	e := http.RegisterRouter(echo.New(), tl)

	e.Logger.Fatal(e.Start(":1323"))
}

func getDbHost() string {
	appEnv := os.Getenv("APP_ENV")
	dbHost := os.Getenv("SQLITE_FILE_PATH")

	if appEnv == "testing" {
		dbHost = ":memory:"
	}

	return dbHost
}
