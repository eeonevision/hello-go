package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eeonevision/hello-go/pkg/repository"
	"github.com/eeonevision/hello-go/pkg/server/handlers/models"
	"github.com/eeonevision/hello-go/pkg/server/handlers/toresponse"
	"github.com/gin-gonic/gin"
)

type handler struct {
	repository repository.API
}

// New initializes new instance of handler.
func New(repository repository.API) *handler {
	return &handler{repository: repository}
}

// PostTodoHandler handles new todo in db.
func (h *handler) PostTodoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			createError(c, err)
			return
		}

		var todo *models.Todo
		if err := json.Unmarshal(body, &todo); err != nil {
			createError(c, err)
			return
		}

		id, err := h.repository.CreateTodo(todo.Title, todo.IsFinished)
		if err != nil {
			createError(c, err)
			return
		}

		success(c, id)
	}
}

// GetAllTodosHandler obtains all todos from repostiry.
func (h *handler) GetAllTodosHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		todos, err := h.repository.FindAllTodos()
		if err != nil {
			createError(c, err)
			return
		}

		success(c, toresponse.ToGetAllTodos(todos))
	}
}

func createError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": "failed",
		"error":  err.Error(),
	})
}

func success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}
