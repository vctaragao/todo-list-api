package functional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/todo-list-api/api/http"
)

func TestCreateTask(t *testing.T) {
	dto := http.TaskDto{TaskId: 0, Description: "Implement functional tests", Priority: 1}

	reqBody, _ := json.Marshal(dto)

	resp := Request("POST", "/create", reqBody)

	assert.Equal(t, 200, resp.StatusCode)

	var task http.TaskDto
	err := DecodeBody(resp, &task)

	assert.NoError(t, err)

	assert.Equal(t, 1, task.TaskId)
	assert.Equal(t, dto.Description, task.Description)
	assert.Equal(t, dto.Priority, task.Priority)
}
