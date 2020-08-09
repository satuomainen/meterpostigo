package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config provides the configuration
var Config Configuration

const configurationFilepath = "meterpostigo.yml"

type Configuration struct {
	Database struct {
		Host     string `yaml:"host" env-required:"true" env:"DB_HOST" env-description:"Database host"`
		Port     string `yaml:"port" env-required:"true" env:"DB_PORT" env-description:"Database port"`
		Username string `yaml:"username" env-required:"true" env:"DB_USER" env-description:"Database user name"`
		Password string `env:"DB_PASSWORD" env-required:"true" env-description:"Database user password"`
		Name     string `yaml:"name" env-required:"true" env:"DB_NAME" env-description:"Database name"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port" env-required:"true" env:"PORT" env-default:"9000"`
	} `yaml:"server"`
}

func Initialize() {
	err := cleanenv.ReadConfig(configurationFilepath, &Config)
	if err != nil {
		panic(fmt.Sprintf("Failed to read configuration '%s', cannot continue - %s", configurationFilepath, err))
	}
}
