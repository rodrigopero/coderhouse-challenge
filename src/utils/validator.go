package utils

import "github.com/go-playground/validator/v10"

var validatorInstance *validator.Validate

func GetValidatorInstance() *validator.Validate {
	if validatorInstance == nil {
		validatorInstance = validator.New(validator.WithRequiredStructEnabled())
	}

	return validatorInstance
}
