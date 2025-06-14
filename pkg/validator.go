package pkg

import (
	"github.com/go-playground/validator/v10"
)

// Validator global
var Validate *validator.Validate

// Inisialisasi validator
func InitValidator() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}
