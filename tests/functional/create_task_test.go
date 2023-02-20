package functional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]interface{}{
		"description": "Milk, eggs, bread",
		"priority":    1,
	})

	resp := Request("POST", "/", reqBody)

	assert.Equal(t, 200, resp.StatusCode)
}
