package create_task

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/todo-list-api/internal/aplication/dto"
	"github.com/vctaragao/todo-list-api/internal/aplication/entity"
	"github.com/vctaragao/todo-list-api/internal/aplication/value_object"
)

type testCase struct {
	id          int
	description string
	priority    int
	expected    int
	err         error
}

type MockRepo struct {
	id  int
	err error
}

func (mr MockRepo) AddTask(t value_object.TaskForCreation) (int, error) {
	if mr.err != nil {
		return 0, mr.err
	}

	return mr.id, nil
}

func (mr MockRepo) DeleteTask(t entity.Task) error {
	return nil
}

var testCases = []testCase{
	{
		description: "isso é uma tarefa",
		id:          1,
		priority:    1,
		expected:    1,
		err:         nil,
	},
	{
		description: "isso é uma tarefa com err",
		id:          1,
		priority:    1,
		expected:    0,
		err:         errors.New("erro vindo da task"),
	},
}

func TestCreateTask(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {

			dto, repo := setup(tc)

			id, err := Create(dto, repo)

			if tc.err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expected, id)
		})
	}
}

func setup(tc testCase) (dto.TaskDto, MockRepo) {
	dto := dto.NewDto(tc.id, tc.description, tc.priority)

	repo := MockRepo{
		id:  tc.expected,
		err: tc.err,
	}

	return dto, repo
}
