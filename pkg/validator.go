package pkg

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type Validator interface {
	// Struct validates the given struct based on defined validation rules.
	// The parameter `s` must be a pointer to a struct.
	// If `s` is invalid, the method will panic.
	// It returns a slice of ValidationErr detailing the validation errors.
	Struct(s interface{}) []ValidationErr
}

type validatorStruct struct {
	validate *validator.Validate
}

func NewValidator() Validator {
	return &validatorStruct{
		validate: validator.New(),
	}
}

// Struct validates the given struct based on defined validation rules.
// The parameter `s` must be a pointer to a struct.
// If `s` is invalid, the method will panic.
// It returns a slice of ValidationErr detailing the validation errors.
func (v *validatorStruct) Struct(s interface{}) []ValidationErr {
	err := v.validate.Struct(s)
	if err == nil {
		return nil
	}

	value := reflect.ValueOf(s)
	if value.Kind() != reflect.Ptr {
		log.Panic().Err(err).Msgf("expected pointer, got %T", s)
	}

	value = value.Elem()
	if value.Kind() != reflect.Struct {
		log.Panic().Err(err).Msgf("expected pointer to struct, got %T", s)
	}

	var errs []ValidationErr
	for _, fe := range err.(validator.ValidationErrors) {
		fieldName := fe.StructField()
		field, ok := value.Type().FieldByName(fieldName)
		if !ok {
			continue
		}

		result := v.parseErr(fe, field)
		if result != nil {
			errs = append(errs, *result)
		}
	}

	return errs
}

func (v *validatorStruct) parseErr(fe validator.FieldError, field reflect.StructField) *ValidationErr {
	tag := fe.Tag()
	fPath, ok := field.Tag.Lookup("json")
	if !ok {
		fPath = field.Name
	}
	fKind := field.Type.Kind().String()

	result := &ValidationErr{
		Tag:  tag,
		Kind: fKind,
		Path: fPath,
	}

	switch tag {
	case "required":
		result.Message = "This field is required"
	case "min":
		result.Message = "This field must be at least " + fe.Param()
	case "max":
		result.Message = "This field must be at most " + fe.Param()
	case "len":
		result.Message = "This field must be " + fe.Param() + " characters long"
	case "email":
		result.Message = "This field must be a valid email address"
	case "oneof":
		result.Options = strings.Split(fe.Param(), " ")
		result.Message = "This field must be one of the options"

	}

	return result
}
