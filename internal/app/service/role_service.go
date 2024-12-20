package service

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/pos/internal/app/contract"
	"github.com/umardev500/pos/internal/app/model"
	"github.com/umardev500/pos/pkg"
)

type roleService struct {
	repo     contract.RoleRepository
	validate pkg.Validator
}

func NewRoleService(r contract.RoleRepository, v pkg.Validator) contract.RoleService {
	return &roleService{
		repo:     r,
		validate: v,
	}
}

func (r *roleService) CreateRole(ctx context.Context, req *model.CreateRoleRequest) *pkg.Response {
	fields := r.validate.Struct(req)
	res := pkg.ValidationResp(fields)
	if res != nil {
		return res
	}

	err := r.repo.CreateRole(ctx, req)
	if err != nil {
		return pkg.DbErrResp(err, req)
	}

	return &pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Role created successfully",
	}
}
