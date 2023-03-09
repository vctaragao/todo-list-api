package storage

import (
	"os"

	"github.com/vctaragao/todo-list-api/internal/application/entity"
	"github.com/vctaragao/todo-list-api/storage/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgress struct {
	Db *gorm.DB
}

func NewPostgress() *Postgress {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		dbPort = "80"
	}

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("unable to connect to database")
	}

	db.AutoMigrate(&schemas.Task{})

	return &Postgress{
		Db: db,
	}
}

func (p *Postgress) AddTask(t entity.Task) (int, error) {
	task := schemas.Task{Description: t.Description, Priority: t.Priority}
	result := p.Db.Create(&task)

	if result.Error != nil {
		return 0, ErrUnableToInsertTask
	}

	return int(task.ID), nil
}
