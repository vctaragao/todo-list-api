package functional

import (
	"bytes"
	"fmt"
	client "net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/api/http"
	"github.com/vctaragao/todo-list-api/internal"
	"github.com/vctaragao/todo-list-api/storage"
)

func Request(method, uri string, reqBody []byte) *client.Response {
	initApp()

	fmt.Println("Building the request")
	req := httptest.NewRequest(method, "http://localhost:1323"+uri, bytes.NewBuffer(reqBody))
	req.RequestURI = ""
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := &client.Client{}
	fmt.Println("Making the request")
	resp, err := client.Do(req)
	fmt.Println("Request Finished")

	if err != nil {
		panic(err)
	}

	return resp
}

func initApp() {
	repo := storage.NewSqLite(":memory:")

	tl := internal.NewTodoList(repo)

	e := http.RegisterRouter(echo.New(), tl)

	go startServer(e)
}

func startServer(e *echo.Echo) {
	go e.Logger.Fatal(e.Start(":1323"))
}
