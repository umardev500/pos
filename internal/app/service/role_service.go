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

func (r *roleService) DeleteRoles(ctx context.Context, req *pkg.IDsReq) *pkg.Response {
	fields := r.validate.Struct(req)
	res := pkg.ValidationResp(fields)
	if res != nil {
		return res
	}

	err := r.repo.DeleteRoles(ctx, req.StringSlice())
	if err != nil {
		return pkg.DbErrResp(err, req)
	}

	return &pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Roles deleted successfully",
	}
}

func (r *roleService) GetRoles(ctx context.Context) *pkg.Response {
	roles, err := r.repo.GetRoles(ctx)
	if err != nil {
		return pkg.DbErrResp(err, nil)
	}

	return &pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Data:       roles,
	}
}

func (r *roleService) GetRoleById(ctx context.Context, id string) *pkg.Response {
	err := r.validate.Uuid(id)
	if err != nil {
		return pkg.AutoSelectErrResp(err)
	}

	role, err := r.repo.GetRoleById(ctx, id)
	if err != nil {
		return pkg.DbErrResp(err, id)
	}

	return &pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Role fetched successfully",
		Data:       role,
	}
}

func (r *roleService) GetRolesByUserId(ctx context.Context, userId string) *pkg.Response {
	err := r.validate.Uuid(userId)
	if err != nil {
		return pkg.AutoSelectErrResp(err)
	}

	roles, err := r.repo.GetRolesByUserId(ctx, userId)
	if err != nil {
		return pkg.DbErrResp(err, userId)
	}

	return &pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Roles fetched successfully",
		Data:       roles,
	}
}

func (r *roleService) UpdateRoleById(ctx context.Context, id string, req *model.UpdateRoleRequest) *pkg.Response {
	err := r.validate.Uuid(id)
	if err != nil {
		return pkg.AutoSelectErrResp(err)
	}

	fields := r.validate.Struct(req)
	res := pkg.ValidationResp(fields)
	if res != nil {
		return res
	}

	err = r.repo.UpdateRoleById(ctx, id, req)
	if err != nil {
		return pkg.DbErrResp(err, req)
	}

	return &pkg.Response{
		StatusCode: fiber.StatusOK,
		Success:    true,
		Message:    "Role updated successfully",
	}
}
