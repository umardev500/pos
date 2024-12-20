package handler

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos/internal/app/contract"
	"github.com/umardev500/pos/internal/app/model"
)

type roleHandler struct {
	service contract.RoleService
}

func NewRoleHandler(s contract.RoleService) contract.RoleHandler {
	return &roleHandler{
		service: s,
	}
}

func (r *roleHandler) CreateRole(c *fiber.Ctx) error {
	var payload model.CreateRoleRequest
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := r.service.CreateRole(ctx, &payload)
	return c.Status(res.StatusCode).JSON(res)
}

func (r *roleHandler) DeleteRoles(c *fiber.Ctx) error {
	return nil
}

func (r *roleHandler) GetRoleById(c *fiber.Ctx) error {
	return nil
}

func (r *roleHandler) GetRolesByUserId(c *fiber.Ctx) error {
	return nil
}

func (r *roleHandler) GetRoles(c *fiber.Ctx) error {
	return nil
}

func (r *roleHandler) UpdateRoleById(c *fiber.Ctx) error {
	return nil
}
