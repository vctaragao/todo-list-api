package main

import (
	"github.com/labstack/echo/v4"
	c "github.com/vctaragao/todo-list-api/src/controllers"
)

func main() {
	e := echo.New()

	e.GET("/", c.HelloWorld)

	e.Logger.Fatal(e.Start(":1323"))
}
