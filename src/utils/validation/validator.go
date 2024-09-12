package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

var validatorInstance *validator.Validate

func GetValidatorInstance() *validator.Validate {
	if validatorInstance == nil {
		validatorInstance = validator.New(validator.WithRequiredStructEnabled())
	}

	return validatorInstance
}

func GetErrors(errors validator.ValidationErrors) string {
	var errorMsgs []string
	var errorMsg string

	for _, fieldError := range errors {
		switch {
		case fieldError.Tag() == "required":
			errorMsg = fmt.Sprintf("the '%s' field is required.", fieldError.Field())
		case fieldError.Tag() == "gte" && fieldError.Kind() == reflect.String:
			errorMsg = fmt.Sprintf("the '%s' field must have at least %s characters.", fieldError.Field(), fieldError.Param())
		case fieldError.Tag() == "gt" && fieldError.Kind() == reflect.String:
			errorMsg = fmt.Sprintf("the '%s' field must have more than %s characters.", fieldError.Field(), fieldError.Param())
		case fieldError.Tag() == "gt" && fieldError.Kind() == reflect.Float64:
			errorMsg = fmt.Sprintf("the '%s' field must be greater than %s.", fieldError.Field(), fieldError.Param())
		case fieldError.Tag() == "gt" && fieldError.Kind() == reflect.Slice:
			errorMsg = fmt.Sprintf("the '%s' field must have more than %s values.", fieldError.Field(), fieldError.Param())
		case fieldError.Tag() == "lte" && fieldError.Kind() == reflect.String:
			errorMsg = fmt.Sprintf("the '%s' field must have less than %s characters.", fieldError.Field(), fieldError.Param())
		case fieldError.Tag() == "alphanum" && fieldError.Kind() == reflect.String:
			errorMsg = fmt.Sprintf("the '%s' field only supports alphanumeric values", fieldError.Field())
		default:
			errorMsg = fmt.Sprintf("the '%s' field have an invalid value.", fieldError.Field())
		}

		errorMsgs = append(errorMsgs, errorMsg)
	}

	return strings.Join(errorMsgs, " ")
}