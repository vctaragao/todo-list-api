package create_task

import (
	"github.com/vctaragao/todo-list-api/internal"
	"github.com/vctaragao/todo-list-api/internal/entity"
)

type Service struct {
	Repo internal.Repository
}

func NewServie(repo internal.Repository) Service {
	return Service{
		Repo: repo,
	}
}

func (s Service) Create(t TaskDto) (int, error) {
	// validation

	// add task to the repo
	task := entity.NewTask(t.Description, t.Priority)

	id, err := s.Repo.AddTask(task)

	if err != nil {
		return 0, err
	}

	return id, nil
}
