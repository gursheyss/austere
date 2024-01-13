package server

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func validURL(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	return strings.Contains(url, "youtube.com") || strings.Contains(url, "youtu.be") || strings.Contains(url, "soundcloud.com")
}

func NewValidator() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("url", validURL)
	return &CustomValidator{validator: v}
}

type ValidationError struct {
	validator.ValidationErrors
}

func (v ValidationError) Error() string {
	for _, err := range v.ValidationErrors {
		switch err.Tag() {
		case "required":
			return fmt.Sprintf("%s is required", err.Field())
		case "url":
			return fmt.Sprintf("%s must be a YouTube or SoundCloud link", err.Field())
		}
	}
	return "Invalid request parameters"
}
