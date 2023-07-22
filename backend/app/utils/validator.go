package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type (
	errorResponse struct {
		FailedField string
		Tag         string
		Value       interface{}
	}

	AddUserRequestValidator struct {
		validator *validator.Validate
	}
)

func (v AddUserRequestValidator) Validate(data interface{}) []errorResponse {
	v.validator = validate

	var validationErrors []errorResponse

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, errorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			})
		}
	}

	return validationErrors
}
