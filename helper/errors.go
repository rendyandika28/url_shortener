package helper

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type httpError struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Errors  []FieldError `json:"errors,omitempty"`
}

type FieldError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	message := err.Error()
	var listErrors []FieldError
	// FIBER / VALIDATOR ERRORS SWITCH
	var fiberError *fiber.Error
	var validationErrors validator.ValidationErrors
	switch {
	case errors.As(err, &fiberError): // Catch for fiber error
		e := err.(*fiber.Error)
		code = e.Code
	case errors.As(err, &validationErrors): // Catch for validation error
		fmt.Println("=====INNNNN======")
		code = fiber.StatusUnprocessableEntity
		for _, fe := range err.(validator.ValidationErrors) {
			listErrors = append(listErrors, FieldError{fe.Field(), msgForTag(fe.Tag(), fe.Field())})
		}
	}

	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		code = fiber.StatusBadRequest
	case errors.Is(err, gorm.ErrRecordNotFound):
		code = fiber.StatusNotFound
	}

	return c.Status(code).JSON(&httpError{
		Status:  code,
		Message: message,
		Errors:  listErrors,
	})
}

func msgForTag(tag string, field string) string {
	tagMessages := map[string]string{
		"required": field + " is required",
		"email":    "Invalid email",
		"url":      "Invalid URL format",
		"slug":     "Invalid slug format",
	}

	return tagMessages[tag]
}
