package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const pathDir = "./configs"

type server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type postgres struct {
	DSN string `yaml:"dsn"`
}

type db struct {
	Postgres postgres `yaml:"postgres"`
}

// Config keeps configuration for service.
type Config struct {
	Debug  bool   `yaml:"debug"`
	Server server `yaml:"server"`
	DB     db     `yaml:"db"`
}

// Parse returns new config from yaml file.
func Parse(fileName string) (*Config, error) {
	var cfg *Config

	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(pathDir)  // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}

	return cfg, viper.Unmarshal(&cfg)
}
