package pkg

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ValidationResp(fields []ValidationErr) *Response {
	if len(fields) == 0 {
		return nil
	}

	refCode := LogError(fmt.Errorf("validation error"))

	return &Response{
		StatusCode: fiber.ErrUnprocessableEntity.Code,
		Message:    "validation error",
		Errors:     fields,
		Code:       "VALIDATION_ERR",
		Ref:        refCode,
	}
}

func AutoSelectErrResp(err error) *Response {
	var resp *Response

	switch err {
	default:
		resp = InServerErrResp(err)
	}

	return resp
}

func InServerErrResp(err error) *Response {
	refCode := LogError(err)
	return &Response{
		StatusCode: fiber.ErrInternalServerError.Code,
		Message:    "internal server error",
		Errors:     nil,
		Code:       "SERVER_ERR",
		Ref:        refCode,
	}
}
