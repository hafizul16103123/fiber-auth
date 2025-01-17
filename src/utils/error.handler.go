package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

// // Error struct is used to represent custom errors with messages
type Error struct {
	Err   string `json:"error"`
	Msg string `json:"message"`
}

func (e *Error) Error() string {
	return e.Msg
}
// FormatValidationError formats the validation errors into a user-friendly message.
func FormatValidationError(err error) string {
	var errorMsgs []string
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			// Build a more user-friendly error message
			errorMsgs = append(errorMsgs, "Field '"+e.Field()+"' failed validation on the '"+e.Tag()+"' tag")
		}
	}
	return strings.Join(errorMsgs, ", ")
}
