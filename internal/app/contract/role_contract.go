package contract

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos/internal/app/model"
	"github.com/umardev500/pos/pkg"
)

type RoleHandler interface {
	CreateRole(c *fiber.Ctx) error
	DeleteRoles(c *fiber.Ctx) error
	GetRoles(c *fiber.Ctx) error
	GetRoleById(c *fiber.Ctx) error
	UpdateRoleById(c *fiber.Ctx) error
}

type RoleService interface {
	CreateRole(ctx context.Context, req *model.CreateRoleRequest) *pkg.Response
	DeleteRoles(ctx context.Context, req *pkg.IDsReq) *pkg.Response
	GetRoles(ctx context.Context) *pkg.Response
	GetRoleById(ctx context.Context, id string) *pkg.Response
	UpdateRoleById(ctx context.Context, id string, req *model.UpdateRoleRequest) *pkg.Response
}

type RoleRepository interface {
	CreateRole(ctx context.Context, data *model.CreateRoleRequest) error
	DeleteRoles(ctx context.Context, ids []string) error
	GetRoles(ctx context.Context) ([]model.Role, error)
	GetRoleById(ctx context.Context, id string) (*model.Role, error)
	UpdateRoleById(ctx context.Context, id string, data *model.UpdateRoleRequest) error
}
