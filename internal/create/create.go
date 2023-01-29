package create

import "github.com/vctaragao/todo-list-api/internal/entity"

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
