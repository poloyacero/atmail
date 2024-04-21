package utils

import (
	"github.com/go-playground/validator/v10"
)

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "email":
		return "Invalid email format " + fe.Param()
	case "min":
		return "Should be greater than " + fe.Param()
	case "validDate":
		return "Invalid date format, should be (YYYY-MM-DD)"
	case "unique":
		return "Not available"
	}

	return "Unknown error"
}
