package entity

import (
	"errors"
	"fmt"
)

const (
	MAX_DESCRIPTION_LENGTH = 140
	MIN_PRIORITY           = 1
	MAX_PRIORITY           = 100
	MIN_ID                 = 1
)

var ErrInvalidId = errors.New("id invalid")
var ErrInvalidDescription = fmt.Errorf("description should be less or egual to %v", MAX_DESCRIPTION_LENGTH)
var ErrInvalidPriority = fmt.Errorf("priority should be in the range of %v-%v", MIN_PRIORITY, MAX_PRIORITY)

type Dto interface {
	GetId() int
	GetDescription() string
	GetPriority() int
}

type Task struct {
	Id          int
	Description string
	Priority    int
}

func NewTask(dto Dto) (Task, error) {
	err := validateDto(dto)

	if err != nil {
		return Task{}, err
	}

	return Task{
		Id:          dto.GetId(),
		Description: dto.GetDescription(),
		Priority:    dto.GetPriority(),
	}, nil
}

func validateDto(dto Dto) (err error) {
	if dto.GetId() < MIN_ID {
		err = ErrInvalidId
	}

	if dto.GetDescription() != "" && len(dto.GetDescription()) > MAX_DESCRIPTION_LENGTH {
		err = ErrInvalidDescription
	}

	if dto.GetPriority() < MIN_PRIORITY || dto.GetPriority() > MAX_PRIORITY {
		err = ErrInvalidPriority
	}

	if err != nil {
		return err
	}

	return nil
}
