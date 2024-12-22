package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos/internal/app/contract"
	"github.com/umardev500/pos/internal/app/model"
	"github.com/umardev500/pos/pkg"
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

	ctx, cancel := pkg.BaseContext()
	defer cancel()

	res := r.service.CreateRole(ctx, &payload)
	return c.Status(res.StatusCode).JSON(res)
}

func (r *roleHandler) DeleteRoles(c *fiber.Ctx) error {
	var payload pkg.IDsReq
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	ctx, cancel := pkg.BaseContext()
	defer cancel()

	ctx = pkg.ParseUnscopedParamContext(c, ctx)

	res := r.service.DeleteRoles(ctx, &payload)
	return c.Status(res.StatusCode).JSON(res)
}

func (r *roleHandler) GetRoleById(c *fiber.Ctx) error {
	var id = c.Params("id")

	ctx, cancel := pkg.ParseUnscopedParamContextBase(c)
	defer cancel()

	res := r.service.GetRoleById(ctx, id)
	return c.Status(res.StatusCode).JSON(res)
}

func (r *roleHandler) GetRoles(c *fiber.Ctx) error {
	ctx, cancel := pkg.ParseUnscopedParamContextBase(c)
	defer cancel()

	res := r.service.GetRoles(ctx)
	return c.Status(res.StatusCode).JSON(res)
}

func (r *roleHandler) UpdateRoleById(c *fiber.Ctx) error {
	var payload model.UpdateRoleRequest
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	ctx, cancel := pkg.ParseUnscopedParamContextBase(c)
	defer cancel()

	var id = c.Params("id")

	res := r.service.UpdateRoleById(ctx, id, &payload)
	return c.Status(res.StatusCode).JSON(res)
}
