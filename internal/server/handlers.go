package server

import (
	"austere/internal/aws"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hey lol")
}

func UploadHandler(c echo.Context) error {
	ctx := context.TODO()
	aws.ConnectToSQS(ctx)
	return c.String(http.StatusOK, "hey lol2")
}
