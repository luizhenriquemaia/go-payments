package utils

import "github.com/go-playground/validator/v10"

type ApiError struct {
	Field string
	Msg   string
}

func get_validation_error_msgs(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "enum":
		return "Invalid value"
	}
	return tag
}

func Get_validation_api_error(errors validator.ValidationErrors) []ApiError {
	api_error := make([]ApiError, len(errors))
	for i, field_error := range errors {
		api_error[i] = ApiError{field_error.Field(), get_validation_error_msgs(field_error.Tag())}
	}
	return api_error
}
