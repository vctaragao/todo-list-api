package create

import (
	"github.com/vctaragao/todo-list-api/internal"
)

type Service struct {
	Repo internal.Repository
}

func CreateService(repo internal.Repository) Service {
	return Service{
		Repo: repo,
	}
}
