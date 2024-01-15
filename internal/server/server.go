package server

import (
	"austere/internal/tools"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	port := tools.EnvPortOr("3002")
	e := echo.New()
	e.GET("/", HelloWorldHandler)
	e.Logger.Fatal(e.Start(port))
}
