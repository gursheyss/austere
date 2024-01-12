package server

import (
	"austere/internal/tools"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	port := tools.EnvPortOr("3001")
	e := echo.New()
	e.GET("/", HelloWorldHandler)
	e.GET("/upload", UploadHandler)
	e.Logger.Fatal(e.Start(port))
}
