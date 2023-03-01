package functional

import (
	"bytes"
	"encoding/json"
	"io"
	net_http "net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/vctaragao/todo-list-api/api/http"
	"github.com/vctaragao/todo-list-api/internal"
	"github.com/vctaragao/todo-list-api/storage"
)

func Request(method, uri string, reqBody []byte) *net_http.Response {
	initApp()
	req := createRequest(method, uri, reqBody)
	client := &net_http.Client{}
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

	go startServer(e)
}

func createRequest(method, uri string, reqBody []byte) *net_http.Request {
	req := httptest.NewRequest(method, "http://localhost:1323"+uri, bytes.NewBuffer(reqBody))
	req.RequestURI = ""
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return req
}

func startServer(e *echo.Echo) {
	go e.Logger.Fatal(e.Start(":1323"))
}

func DecodeBody(resp *net_http.Response, dto interface{}) error {
	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &dto)

	if err != nil {
		return err
	}

	return nil
}
