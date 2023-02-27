package functional

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/todo-list-api/api/http"
)

func TestCreateTask(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]interface{}{
		"description": "Milk, eggs, bread",
		"priority":    1,
	})

	resp := Request("POST", "/create", reqBody)

	assert.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	assert.NoError(t, err)
	fmt.Printf("%v", string(body))

	var tasks []http.TaskDto
	err = json.Unmarshal(body, &tasks)

	assert.NoError(t, err)

	fmt.Printf("%v", tasks)
}
