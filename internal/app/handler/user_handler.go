package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos/internal/app/contract"
	"github.com/umardev500/pos/pkg"
)

type userHandler struct {
	s contract.UserService
}

func NewUserHandler(s contract.UserService) contract.UserHandler {
	return &userHandler{
		s: s,
	}
}

func (u *userHandler) GetUsers(c *fiber.Ctx) error {
	ctx, cancel := pkg.ParseUnscopedParamContextBase(c)
	defer cancel()

	res := u.s.GetUsers(ctx)
	return c.Status(res.StatusCode).JSON(res)
}
