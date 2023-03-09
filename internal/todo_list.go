package internal

import (
	"github.com/vctaragao/todo-list-api/internal/aplication"
	"github.com/vctaragao/todo-list-api/internal/aplication/create_task"
	"github.com/vctaragao/todo-list-api/internal/aplication/dto"
)

type TodoList struct {
	Repo aplication.Repository
}

func NewTodoList(repo aplication.Repository) TodoList {
	return TodoList{
		Repo: repo,
	}
}

func (tl TodoList) CreateTask(description string, priority int) (int, error) {
	dto := dto.NewDto(0, description, priority)

	return create_task.Create(dto, tl.Repo)
}
