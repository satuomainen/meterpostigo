package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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
	CORS struct {
		// See https://echo.labstack.com/middleware/cors
		AllowOrigins []string `yaml:"allowOrigins"`
	} `yaml:"cors"`
}

func GetCORSConfig() middleware.CORSConfig {

	var allowOrigins []string = nil

	if Config.CORS.AllowOrigins != nil && len(Config.CORS.AllowOrigins) > 0 {
		allowOrigins = Config.CORS.AllowOrigins
	}

	log.Infof("Configured AllowOrigins: %v", allowOrigins)

	var corsConfig = middleware.CORSConfig{
		Skipper:          nil,
		AllowOrigins:     allowOrigins,
		AllowMethods:     nil,
		AllowHeaders:     nil,
		AllowCredentials: false,
		ExposeHeaders:    nil,
		MaxAge:           0,
	}

	return corsConfig
}

func Initialize() {
	err := cleanenv.ReadConfig(configurationFilepath, &Config)
	if err != nil {
		panic(fmt.Sprintf("Failed to read configuration '%s', cannot continue - %s", configurationFilepath, err))
	}
}
