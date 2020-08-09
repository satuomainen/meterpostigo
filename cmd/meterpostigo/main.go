package main

import (
	"com.github/satuomainen/meterpostigo/internal/api"
	"com.github/satuomainen/meterpostigo/internal/config"
	"com.github/satuomainen/meterpostigo/internal/db"
	"com.github/satuomainen/meterpostigo/internal/server"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Initialize()
	db.Initialize()

	e := echo.New()

	e.Use(middleware.Logger())

	var server server.MetricsServer
	serverapi.RegisterHandlers(e, &server)

	address := fmt.Sprintf("0.0.0.0:%s", config.Config.Server.Port)
	e.Logger.Fatal(e.Start(address))
}
