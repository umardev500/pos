package pkg

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

func DbErrResp(err error, model interface{}) *Response {
	var resp *Response

	switch err := err.(type) {
	case *pgconn.PgError:
		resp = parsePgErr(err, model)
	}

	return resp
}

func parsePgErr(err *pgconn.PgError, model interface{}) *Response {
	if err.Code == "23505" {
		return parseUniqueConstraint(err, model)
	}

	return InServerErrResp(err)
}

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
