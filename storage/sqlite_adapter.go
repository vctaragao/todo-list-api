package storage

import (
	"github.com/vctaragao/todo-list-api/internal/application/entity"
	"github.com/vctaragao/todo-list-api/storage/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqLite struct {
	Db *gorm.DB
}

func NewSqLite(file string) *SqLite {

	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})

	if err != nil {
		panic("unable to connect to database")
	}

	db.AutoMigrate(&schemas.Task{})

	return &SqLite{
		Db: db,
	}
}

func (s *SqLite) AddTask(t entity.Task) (int, error) {
	task := schemas.Task{Description: t.Description, Priority: t.Priority}
	result := s.Db.Create(&task)

	if result.Error != nil {
		return 0, ErrUnableToInsertTask
	}

	return int(task.ID), nil
}
