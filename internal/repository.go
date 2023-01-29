package internal

import "github.com/vctaragao/todo-list-api/internal/entity"

type Repository interface {
	AddTask(t entity.Task) (int, error)
}
