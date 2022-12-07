package validator

import "github.com/go-playground/validator/v10"

var instance *validator.Validate

// GetValidator will return a validator instance
func GetValidator() *validator.Validate {
	if instance == nil {
		instance = validator.New()
	}
	return instance
}
