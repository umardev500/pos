package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos/internal/app/contract"
	"github.com/umardev500/pos/pkg"
)

type userService struct {
	r contract.UserRepository
	v pkg.Validator
}

func NewUserService(r contract.UserRepository, v pkg.Validator) contract.UserService {
	return &userService{
		r: r,
		v: v,
	}
}

func (r *userService) GetUsers(ctx context.Context) *pkg.Response {
	users, err := r.r.GetUsers(ctx)
	if err != nil {
		return pkg.DbErrResp(err, nil)
	}

	return &pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Data:       users,
	}
}
