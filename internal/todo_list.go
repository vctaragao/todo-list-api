package internal

import (
	"github.com/vctaragao/todo-list-api/internal/application"
	"github.com/vctaragao/todo-list-api/internal/application/create_task"
)

type TodoList struct {
	Repo application.Repository
}

func NewTodoList(repo application.Repository) TodoList {
	return TodoList{
		Repo: repo,
	}
}

func (tl TodoList) CreateTask(description string, priority int) (int, error) {
	dto := create_task.NewDto(description, priority)
	return create_task.Create(dto, tl.Repo)
}
