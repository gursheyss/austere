package server

import (
	"austere/ytdlp"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type BodyParams struct {
	URL    string `form:"url" validate:"required,url"`
	Title  string `form:"title"`
	Album  string `form:"album"`
	Artist string `form:"artist"`
}

func HelloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hey lol")
}

func UploadHandler(c echo.Context) error {
	params := new(BodyParams)
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
	ytdlp.Download()
	return c.String(http.StatusOK, "hey lol2")
}
