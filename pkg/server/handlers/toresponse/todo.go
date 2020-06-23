package toresponse

import (
	repositoryModels "github.com/eeonevision/hello-go/pkg/repository/models"
	"github.com/eeonevision/hello-go/pkg/server/handlers/models"
)

// ToGetAllTodos converts repository's model to handler.
func ToGetAllTodos(in []*repositoryModels.Todo) []*models.Todo {
	todos := make([]*models.Todo, len(in))

	for i, todo := range in {
		todos[i] = &models.Todo{
			ID:         int(todo.ID),
			Title:      todo.Title,
			IsFinished: todo.IsFinished,
		}
	}

	return todos
}
