package contract

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos/internal/app/model"
	"github.com/umardev500/pos/pkg"
)

type UserHandler interface {
	// CreateUser(c *fiber.Ctx) error
	// DeleteUsers(c *fiber.Ctx) error
	// GetUserById(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	// UpdateUserById(c *fiber.Ctx) error
}

type UserService interface {
	GetUsers(ctx context.Context) *pkg.Response
}

type UserRepository interface {
	// CreateUser(ctx context.Context, data *model.CreateUserReq) error
	// DeleteUsers(ctx context.Context, ids []string) error
	// GetUserById(ctx context.Context, id string) (model.User, error)
	GetUsers(ctx context.Context) ([]model.User, error)
}
