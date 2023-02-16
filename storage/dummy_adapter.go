package storage

import (
	"math/rand"

	"github.com/vctaragao/todo-list-api/internal/application/entity"
)

type DummyAdapter struct {
}

func (repo DummyAdapter) AddTask(t entity.Task) (int, error) {
	return rand.Intn(100), nil
}

func NewDummyAdapter() DummyAdapter {
	return DummyAdapter{}
}
