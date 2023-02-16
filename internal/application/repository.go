package application

import "github.com/vctaragao/todo-list-api/internal/application/entity"

type Repository interface {
	AddTask(t entity.Task) (int, error)
}
