package helper

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidationSlug(fl validator.FieldLevel) bool {
	slug := fl.Field().String()

	// Use a regular expression to check if the string is a valid slug
	// The pattern allows lowercase alphanumeric characters, hyphens, and underscores
	match, _ := regexp.MatchString("^[a-zA-Z0-9_\\-]+$", slug)

	return match
}
