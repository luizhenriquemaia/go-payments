package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func get_validation_error_msgs(tag string, err_param *string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "enum":
		return "Invalid value"
	case "min_length":
		return "Invalid field length, this field requires at least " + *err_param + " characters"
	case "equal_length":
		return "Invalid field length, this field requires " + *err_param + " characters"
	case "max_length":
		return "Invalid field length, this field requires less than " + *err_param + " characters"
	case "only_digits":
		return "This field must contain only digits"
	}
	return tag
}

func Get_validation_api_error(errors validator.ValidationErrors) []ApiError {
	api_error := make([]ApiError, len(errors))
	for i, field_error := range errors {
		details := field_error.Param()
		api_error[i] = ApiError{
			strings.ToLower(field_error.Field()),
			get_validation_error_msgs(field_error.Tag(), &details),
		}
	}
	return api_error
}
