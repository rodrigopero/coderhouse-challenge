package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validatorInstance *validator.Validate

func GetValidatorInstance() *validator.Validate {
	if validatorInstance == nil {
		validatorInstance = validator.New(validator.WithRequiredStructEnabled())
	}

	return validatorInstance
}

func GetErrorList(errors validator.ValidationErrors) []string {
	var errorMsgs []string
	var errorMsg string

	for _, fieldError := range errors {

		switch fieldError.Tag() {
		case "required":
			errorMsg = fmt.Sprintf("The '%s' field is required.", fieldError.Field())
		case "gte":
			errorMsg = fmt.Sprintf("The '%s' field must have %s characters at least.", fieldError.Field(), fieldError.Param())
		case "gt":
			errorMsg = fmt.Sprintf("The '%s' field must be greater than %s.", fieldError.Field(), fieldError.Param())
		case "lte":
			errorMsg = fmt.Sprintf("The '%s' field must have a maximum of %s characters", fieldError.Field(), fieldError.Param())
		case "alphanum":
			errorMsg = fmt.Sprintf("The '%s' field only supports alphanumeric values", fieldError.Field())
		default:
			errorMsg = fmt.Sprintf("The '%s' field have an invalid value.", fieldError.Field())
		}

		errorMsgs = append(errorMsgs, errorMsg)

	}

	return errorMsgs
}
