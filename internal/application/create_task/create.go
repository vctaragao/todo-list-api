package create_task

import (
	"fmt"

	"github.com/vctaragao/todo-list-api/internal/application"
	"github.com/vctaragao/todo-list-api/internal/application/entity"
)

var ErrUnableToCreateTask = fmt.Errorf("unable to create task")

func Create(t TaskDto, repo application.Repository) (int, error) {
	task := entity.NewTask(t.Description, t.Priority)

	id, err := repo.AddTask(task)

	if err != nil {
		return 0, ErrUnableToCreateTask
	}

	return id, nil
}
