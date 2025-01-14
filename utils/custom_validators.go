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

func validateEnum(field_level validator.FieldLevel) bool {
	value, ok := field_level.Field().Interface().(Enum)
	if !ok {
		log.Printf("fail on converting Enum %v", field_level)
		panic("failed the conversion to enum in validation")
	}
	return value.IsValid()
}

func validateOnlyDigits(field_level validator.FieldLevel) bool {
	request_value, ok := field_level.Field().Interface().(string)
	if !ok {
		log.Printf("fail on converting to string the value %v", field_level.Field().Interface())
		panic("failed the conversion to string in only digits validation")
	}
	re := regexp.MustCompile(`\D+`)
	parsed_str := re.ReplaceAllString(request_value, "")
	parsed_length := len(parsed_str)
	if parsed_length > 0 {
		field_level.Field().SetString(parsed_str)
		return true
	}
	return false
}

func validateEqualLength(field_level validator.FieldLevel) bool {
	request_value := field_level.Field().Interface().(string)
	validation_length, err := strconv.ParseInt(field_level.Param(), 10, 0)
	if err != nil {
		panic("Equal length validation number couldn't be converted into a integer!")
	}
	if int(validation_length) == len(request_value) {
		return true
	}
	return false
}

func validateMinLength(field_level validator.FieldLevel) bool {
	request_value := field_level.Field().Interface().(string)
	min_length, err := strconv.ParseInt(field_level.Param(), 10, 0)
	if err != nil {
		panic("Minimum length validation number couldn't be converted into a integer!")
	}
	if len(request_value) >= int(min_length) {
		return true
	}
	return false
}

func validateMaxLength(field_level validator.FieldLevel) bool {
	request_value := field_level.Field().Interface().(string)
	max_length, err := strconv.ParseInt(field_level.Param(), 10, 0)
	if err != nil {
		panic("Max length validation number couldn't be converted into a integer!")
	}
	if len(request_value) <= int(max_length) {
		return true
	}
	return false
}

func InitCustomValidators() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("enum", validateEnum)
	}
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("only_digits", validateOnlyDigits)
	}
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("equal_length", validateEqualLength)
	}
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("min_length", validateMinLength)
	}
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("max_length", validateMaxLength)
	}
	log.Print("custom validators loaded")
}
