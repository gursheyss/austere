package server

import (
	"austere/internal/models"
	"austere/internal/ytdlp"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func HelloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hey lol")
}

func UploadHandler(c echo.Context) error {
	params := new(models.BodyParams)
	if err := c.Bind(params); err != nil {
		return err
	}
	if err := c.Validate(params); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return c.String(http.StatusBadRequest, ValidationError{errs}.Error())
		}
		return err
	}
	fmt.Println(params)
	ytdlp.Download(params)
	return c.String(http.StatusOK, "hey lol2")
}
