package validator

import (
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
	return v.Struct(data)
}
