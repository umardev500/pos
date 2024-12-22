package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/pos/internal/app/contract"
	"github.com/umardev500/pos/internal/app/model"
	"github.com/umardev500/pos/pkg"
)

type roleRepository struct {
	db *pkg.GormInstance
}

func NewRoleRepository(db *pkg.GormInstance) contract.RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) CreateRole(ctx context.Context, data *model.CreateRoleRequest) error {
	data.ID = uuid.New()
	return r.db.GetConn(ctx).Create(data).Error
}

func (r *roleRepository) DeleteRoles(ctx context.Context, ids []string) error {
	var conn = r.db.GetConn(ctx)

	return conn.Where("id IN ?", ids).Delete(&model.Role{}).Error
}

func (r *roleRepository) GetRoleById(ctx context.Context, id string) (*model.Role, error) {
	var role model.Role
	var conn = r.db.GetConn(ctx)

	err := conn.First(&role, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) GetRoles(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	var conn = r.db.GetConn(ctx)

	err := conn.Find(&roles).Error

	return roles, err
}

func (r *roleRepository) UpdateRoleById(ctx context.Context, id string, data *model.UpdateRoleRequest) error {
	var conn = r.db.GetConn(ctx)

	return conn.Model(&model.Role{}).Update("name", data.Name).Error
}
