package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type validatorS struct {
	*validator.Validate
}

type Validator interface {
	ValidateStruct(data interface{}) error
}

func NewValidator() Validator {
	return validatorS{validator.New()}
}

func (v validatorS) ValidateStruct(data interface{}) error {
	err := v.Struct(data)
	if err == nil {
		return nil
	}

	errors := ""

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		switch err.Tag() {
		case "required":
			errors += fmt.Sprintf("Field %s is required\n", fieldName)
		case "email":

		default:
			errors += fmt.Sprintf("Field %s is not valid: %s\n", fieldName, err.Tag())
		}
	}
	return fmt.Errorf(errors)

}
