package server

import (
	"fmt"
	"net/http"

	"github.com/eeonevision/hello-go/internal/config"
	"github.com/eeonevision/hello-go/pkg/repository"
	"github.com/eeonevision/hello-go/pkg/server/handlers"
	"github.com/gin-gonic/gin"
)

// Run starts new server.
func Run(cfg *config.Config, repository repository.API) {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	h := handlers.New(repository)

	v1 := r.Group("v1")
	v1.POST("/todos", h.PostTodoHandler())
	v1.GET("/todos", h.GetAllTodosHandler())

	r.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
}
