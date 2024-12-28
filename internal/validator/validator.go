package validator

import (
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
