package repository

import "github.com/eeonevision/hello-go/pkg/repository/models"

// API is public interface for repository methods.
type API interface {
	CreateTodo(title string, isFinished bool) (int, error)
	FindAllTodos() ([]*models.Todo, error)
}
