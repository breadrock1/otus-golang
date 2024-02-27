package hw09structvalidator

import "reflect"

// tag to define rules of validation.
const validatorTag = "validate"

// structure validation rules array.
type structRules []fieldRules

// validation rules for single field.
type fieldRules interface {
	fieldName() string
	validate(errs ValidationErrors, value reflect.Value) ValidationErrors
}

// field type to validate.
type validateKind int

const (
	validateRegular validateKind = iota
	validateSlice
)
