package validations

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GenericApiErrorValidator(err error) []ValidationError {
	var errors []ValidationError

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			message := getValidationMessage(fieldErr)
			errors = append(errors, ValidationError{
				Field:   fieldErr.Field(),
				Message: message,
			})
		}
	} else {
		errors = append(errors, ValidationError{
			Field:   "RequestBody",
			Message: "Invalid Request Body",
		})
	}
	return errors
}

func getValidationMessage(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "required":
		return fmt.Sprintf("%s is required",
			fieldErr.Field(),
		)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long",
			fieldErr.Field(),
			fieldErr.Param(),
		)
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long",
			fieldErr.Field(),
			fieldErr.Param(),
		)
	case "email":
		return fmt.Sprintf("%s must be a valid email address",
			fieldErr.Field(),
		)
	case "eqfield":
		return fmt.Sprintf("%s must be equal to %s",
			fieldErr.Field(),
			fieldErr.Param(),
		)
	default:
		return fmt.Sprintf("%s is not valid",
			fieldErr.Field(),
		)
	}
}
