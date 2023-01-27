package main

import (
	"github.com/labstack/echo/v4"
	h "github.com/vctaragao/todo-list-api/internal/http"
)

func main() {
	e := echo.New()

	e.GET("/", h.HelloWorld)

	e.Logger.Fatal(e.Start(":1323"))
}
