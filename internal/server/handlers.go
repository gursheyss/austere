package server

import (
	"austere/internal/models"
	"austere/internal/prometheus"
	"austere/internal/ytdlp"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func HelloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hey lol")
}

func UploadHandler(c echo.Context) error {
    startTime := time.Now()
    prometheus.UploadsCounter.Inc()

    params := make([]models.BodyParams, 0)
    if err := c.Bind(&params); err != nil {
        prometheus.FailedUploadsCounter.Inc()
        return err
    }

    if len(params) == 0 {
        prometheus.FailedUploadsCounter.Inc()
        return c.String(http.StatusBadRequest, "No params provided")
    }

    v := validator.New()
    for _, param := range params {
    if err := v.Struct(param); err != nil {
        if errs, ok := err.(validator.ValidationErrors); ok {
            prometheus.FailedUploadsCounter.Inc()
            return c.String(http.StatusBadRequest, ValidationError{errs}.Error())
        }
        prometheus.FailedUploadsCounter.Inc()
        return err
    }
    fmt.Println(param)
    err := ytdlp.Download(&param)
    if err != nil {
        prometheus.FailedUploadsCounter.Inc()
        return err
    }
}

    prometheus.UploadDuration.Observe(time.Since(startTime).Seconds())
    return c.String(http.StatusOK, "hey lol2")
}