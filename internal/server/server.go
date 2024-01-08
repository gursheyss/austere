package server

import (
	"austere/internal/tools"
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	port := tools.EnvPortOr("3000")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(port))
}
