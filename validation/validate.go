// Package validation provides types and methods for performing validation
package validation

import validator "gopkg.in/go-playground/validator.v9"

// Validate exposes a reusable validator value
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}
