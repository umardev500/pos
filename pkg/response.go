package pkg

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func ValidationResp(fields []ValidationErr) *Response {
	if len(fields) == 0 {
		return nil
	}

	refCode := LogError(fmt.Errorf("validation error"))

	return &Response{
		StatusCode: fiber.ErrUnprocessableEntity.Code,
		Message:    "Validation error",
		Errors:     fields,
		Code:       "VALIDATION_ERR",
		Ref:        refCode,
	}
}

func AutoSelectErrResp(err error) *Response {
	var resp *Response

	switch err {
	case ErrInvalidId:
		resp = InvalidIdResp(err)
	case gorm.ErrRecordNotFound:
		resp = NotFondResp(err)
	default:
		resp = InServerErrResp(err)
	}

	return resp
}

func InvalidIdResp(err error) *Response {
	errCode := LogError(err)
	return &Response{
		StatusCode: fiber.ErrUnprocessableEntity.Code,
		Message:    "Invalid ID format. The provided ID is not a valid UUID",
		Errors:     nil,
		Code:       "INVALID_ID",
		Ref:        errCode,
	}
}

// InServerErrResp generates a response for internal server errors
//
// Parameters:
// - err: The error to be logged.
//
// Returns:
// - A Response object for internal server errors.
func InServerErrResp(err error) *Response {
	refCode := LogError(err)
	return &Response{
		StatusCode: fiber.ErrInternalServerError.Code,
		Message:    "Internal server error",
		Errors:     nil,
		Code:       "SERVER_ERR",
		Ref:        refCode,
	}
}

func NotFondResp(err error) *Response {
	refCode := LogError(err)
	return &Response{
		StatusCode: fiber.StatusNotFound,
		Message:    "Resource not found",
		Errors:     nil,
		Code:       "NOT_FOUND",
		Ref:        refCode,
	}
}

// DbErrResp proccess the given error and generates an appropriate response based on the error type
//
// Parameters:
// - err: The PostgreSQL error to process.
// - model: A pointer to the struct containing the request input data used for error response generation.
//
// Returns:
// - A Response object based on the error type.
func DbErrResp(err error, model interface{}) *Response {
	var resp *Response

	switch err := err.(type) {
	case *pgconn.PgError:
		resp = parsePgErr(err, model)
	default:
		resp = AutoSelectErrResp(err)
	}

	return resp
}

// parsePgErr handles PostgreSQL errors and returns an appropriate response.
//
// Parameters:
// - err: The PostgreSQL error to process.
// - model: A pointer to the struct containing the request input data used for error response generation.
//
// Returns:
// - A Response object based on the error type.
func parsePgErr(err *pgconn.PgError, model interface{}) *Response {
	if err.Code == "23505" {
		return parseUniqueConstraint(err, model)
	}

	return InServerErrResp(err)
}

// parseUniqueConstraint parse unique constraint
//
// Parameters:
// - err: The PostgreSQL error to process.
// - model: A pointer to the struct containing the request input data used for error response generation.
//
// Returns:
// - A Response object based on the error type.
func parseUniqueConstraint(err *pgconn.PgError, model interface{}) *Response {
	constraintName := err.ConstraintName

	reflectType := reflect.TypeOf(model)
	if reflectType.Kind() != reflect.Ptr {
		log.Panic().Msgf("expected pointer, got %T", model)
	}

	if reflectType.Elem().Kind() != reflect.Struct {
		log.Panic().Msgf("expected pointer to struct, got %T", model)
	}
	reflectType = reflectType.Elem()

	var fe *ValidationErr

	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		if tagConstraint, ok := field.Tag.Lookup("unique_contraint"); ok {
			if constraintName == tagConstraint {
				fPath, ok := field.Tag.Lookup("json")
				if !ok {
					fPath = field.Name
				}

				fe = &ValidationErr{
					Tag:     "unique",
					Kind:    field.Type.Kind().String(),
					Path:    fPath,
					Message: "This field must be unique",
				}

				break
			}
		}
	}

	return &Response{
		StatusCode: fiber.ErrBadRequest.Code,
		Message:    "Constraint violation error",
		Code:       string(ErrUniqueConstraint),
		Errors:     fe,
	}
}
