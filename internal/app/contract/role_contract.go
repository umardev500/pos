package contract

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/pos/internal/app/model"
	"github.com/umardev500/pos/pkg"
)

type RoleHandler interface {
	CreateRole(c *fiber.Ctx) error
	DeleteRoles(c *fiber.Ctx) error
	GetRoles(c *fiber.Ctx) error
	GetRoleById(c *fiber.Ctx) error
	GetRolesByUserId(c *fiber.Ctx) error
	UpdateRoleById(c *fiber.Ctx) error
}

type RoleService interface {
	CreateRole(ctx context.Context, req *model.CreateRoleRequest) *pkg.Response
}

type RoleRepository interface {
	CreateRole(ctx context.Context, data *model.CreateRoleRequest) error
	DeleteRoles(ctx context.Context, ids []uuid.UUID) error
	GetRoles(ctx context.Context) ([]model.Role, error)
	GetRoleById(ctx context.Context, id uuid.UUID) (*model.Role, error)
	GetRolesByUserId(ctx context.Context, userId uuid.UUID) ([]model.Role, error)
	UpdateRoleById(ctx context.Context, id uuid.UUID, data *model.UpdateRoleRequest) error
}
