package main

import (
	"os"

	"github.com/eeonevision/hello-go/internal/config"
	"github.com/eeonevision/hello-go/pkg/repository/postgres"
	"github.com/eeonevision/hello-go/pkg/server"
)

const (
	configNameEnv     = "HELLO_GO_FILE_NAME"
	defaultConfigName = "dev.yml"
)

func main() {
	fileName := os.Getenv(configNameEnv)

	if fileName == "" {
		fileName = defaultConfigName
	}

	cfg, err := config.Parse(fileName)
	if err != nil {
		panic(err)
	}

	storage := postgres.New(cfg)

	server.Run(cfg, storage)
}
