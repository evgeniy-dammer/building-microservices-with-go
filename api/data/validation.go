package data

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidationError wraps the validators FieldError
type ValidationError struct {
	validator.FieldError
}

// Error returns the error strings
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation contains
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		return nil
	}

	return &Validation{validate}
}

// Validate validates a struct after deserializing JSON
func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		// cast the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// validateSKU is custom function for SKU validation
func validateSKU(fl validator.FieldLevel) bool {
	// string format
	reg := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")

	// searching a string with format below
	matches := reg.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}
