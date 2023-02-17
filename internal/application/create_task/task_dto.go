package create_task

import (
	"fmt"
)

const (
	MAX_DESCRIPTION_LENGTH = 140
	MIN_PRIORITY           = 1
	MAX_PRIORITY           = 100
)

var ErrInvalidDescription = fmt.Errorf("description should be less or egual to %v", MAX_DESCRIPTION_LENGTH)
var ErrInvalidPriority = fmt.Errorf("priority should be in the range of %v-%v", MIN_PRIORITY, MAX_PRIORITY)

type TaskDto struct {
	Description string
	Priority    int
}

func NewDto(description string, priority int) (dto TaskDto, err error) {

	if len(description) > MAX_DESCRIPTION_LENGTH {
		err = ErrInvalidDescription
	}

	if priority < MIN_PRIORITY || priority > MAX_PRIORITY {
		err = ErrInvalidPriority
	}

	if err != nil {
		return dto, err
	}

	dto = TaskDto{
		Description: description,
		Priority:    priority,
	}

	return dto, nil
}
