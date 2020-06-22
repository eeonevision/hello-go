package main

import (
	"os"

	"github.com/eeonevision/hello-go/internal/config"
	"github.com/eeonevision/hello-go/pkg/repository/postgres"
	"github.com/eeonevision/hello-go/pkg/server"
)

const fileNameEnv = "HELLO_GO_FILE_NAME"

func main() {
	fileName := os.Getenv(fileNameEnv)

	cfg, err := config.Parse(fileName)
	if err != nil {
		panic(err)
	}

	storage := postgres.New(cfg)

	server.Run(cfg, storage)
}
