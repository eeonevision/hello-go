package postgres

import (
	"github.com/eeonevision/hello-go/internal/config"
	"github.com/eeonevision/hello-go/pkg/repository/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type storage struct {
	db *gorm.DB
}

func New(cfg *config.Config) *storage {
	db, err := gorm.Open("postgres", cfg.DB.Postgres.DSN)
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Todo{})

	return &storage{db: db}
}

func (s *storage) CreateTodo(title string, isFinished bool) (id int, err error) {
	todo := &models.Todo{
		Title:      title,
		IsFinished: isFinished,
	}

	if err = s.db.Create(todo).Error; err != nil {
		return
	}

	id = int(todo.ID)

	return
}

func (s *storage) FindAllTodos() (todos []*models.Todo, err error) {
	return todos, s.db.Model(&models.Todo{}).Find(&todos).Error
}
