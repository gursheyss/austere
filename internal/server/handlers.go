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
	params := make([]models.BodyParams, 0)
	if err := c.Bind(&params); err != nil {
		return err
	}

	v := validator.New()
	for _, param := range params {
		if err := v.Struct(param); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				return c.String(http.StatusBadRequest, ValidationError{errs}.Error())
			}
			return err
		}
		fmt.Println(param)
		ytdlp.Download(&param)
	}
	return c.String(http.StatusOK, "hey lol2")
}
