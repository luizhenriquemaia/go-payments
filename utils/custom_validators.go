package utils

import (
	"log"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Enum interface {
	IsValid() bool
}

func Validate_enum(field_level validator.FieldLevel) bool {
	value := field_level.Field().Interface().(Enum)
	return value.IsValid()
}

func Validate_only_digits(field_level validator.FieldLevel) bool {
	request_value := field_level.Field().Interface().(string)
	re := regexp.MustCompile(`\D+`)
	parsed_str := re.ReplaceAllString(request_value, "")
	parsed_length := len(parsed_str)
	if parsed_length > 0 {
		field_level.Field().SetString(parsed_str)
		return true
	}
	return false
}

func Validate_equal_length(field_level validator.FieldLevel) bool {
	request_value := field_level.Field().Interface().(string)
	validation_length, err := strconv.ParseInt(field_level.Param(), 10, 0)
	if err != nil {
		panic("Final length of validate only digits couldn't be converted into a integer!")
	}
	if int(validation_length) == len(request_value) {
		return true
	}
	return false
}

func Init_custom_validators() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("enum", Validate_enum)
	}
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("only_digits", Validate_only_digits)
	}
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("equal_length", Validate_equal_length)
	}
	log.Print("custom validators loaded")
}
