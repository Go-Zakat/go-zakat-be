package validator

import "github.com/go-playground/validator/v10"

// NewValidator mengembalikan instance validator
func NewValidator() *validator.Validate {
	v := validator.New()

	// Di sini kamu bisa tambahin custom validation
	// contoh: v.RegisterValidation("password", validatePasswordStrength)

	return v
}
