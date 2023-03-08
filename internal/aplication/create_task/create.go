package create_task

import (
	"fmt"

	"github.com/vctaragao/todo-list-api/internal/aplication"
	"github.com/vctaragao/todo-list-api/internal/aplication/dto"
	"github.com/vctaragao/todo-list-api/internal/aplication/value_object"
)

var ErrUnableToCreateTask = fmt.Errorf("unable to create task")

func Create(dto dto.TaskDto, repo aplication.Repository) (int, error) {
	task := value_object.NewTaskForCreation(dto.Description, dto.Priority)

	id, err := repo.AddTask(task)

	if err != nil {
		return 0, ErrUnableToCreateTask
	}

	return id, nil
}
