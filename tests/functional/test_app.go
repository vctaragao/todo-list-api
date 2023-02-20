package functional

import (
	"bytes"
	client "net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/api/http"
	"github.com/vctaragao/todo-list-api/internal"
	"github.com/vctaragao/todo-list-api/storage"
)

func Request(method, uri string, reqBody []byte) *client.Response {
	go initApp()

	req := httptest.NewRequest(method, "http://localhost:1323"+uri, bytes.NewBuffer(reqBody))
	req.RequestURI = ""
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	client := &client.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return resp
}

func initApp() {
	repo := storage.NewSqLite(":memory:")

	tl := internal.NewTodoList(repo)

	e := http.RegisterRouter(echo.New(), tl)

	e.Logger.Fatal(e.Start(":1323"))
}
