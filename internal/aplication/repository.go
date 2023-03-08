package aplication

import (
	"github.com/vctaragao/todo-list-api/internal/aplication/entity"
	"github.com/vctaragao/todo-list-api/internal/aplication/value_object"
)

type Repository interface {
	AddTask(t value_object.TaskForCreation) (int, error)
	DeleteTask(t entity.Task) error
}
