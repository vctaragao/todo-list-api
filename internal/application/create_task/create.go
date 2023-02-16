package create_task

import (
	"github.com/vctaragao/todo-list-api/internal/application"
	"github.com/vctaragao/todo-list-api/internal/application/entity"
)

func Create(t TaskDto, repo application.Repository) (int, error) {
	// validation

	// add task to the repo
	task := entity.NewTask(t.Description, t.Priority)

	id, err := repo.AddTask(task)

	if err != nil {
		return 0, err
	}

	return id, nil
}
