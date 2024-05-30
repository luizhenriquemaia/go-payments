package utils

import (
	"log"

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

func Init_custom_validators() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("enum", Validate_enum)
	}
	log.Print("custom validators loaded")
}
