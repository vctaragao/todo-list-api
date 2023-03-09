package delete_task

import (
	"errors"

	"github.com/vctaragao/todo-list-api/internal/aplication"
	"github.com/vctaragao/todo-list-api/internal/aplication/dto"
	"github.com/vctaragao/todo-list-api/internal/aplication/entity"
)

var ErrUnableToDeleteTask error = errors.New("unable to remove task")

func Delete(dto dto.TaskDto, repo aplication.Repository) error {
	task, err := entity.NewTask(dto)
	if err != nil {
		return err
	}

	err = repo.DeleteTask(task)

	if err != nil {
		return ErrUnableToDeleteTask
	}

	return nil
}
