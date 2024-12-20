package model

import (
	"github.com/google/uuid"
	"github.com/umardev500/pos/pkg"
)

type Role struct {
	ID   uuid.UUID `json:"id" gorm:"primaryKey"`
	Name string    `json:"name"`
	pkg.TimeCommon
}

func (Role) TableName() string {
	return "roles"
}

type CreateRoleRequest struct {
	ID   uuid.UUID `json:"-"`
	Name string    `json:"name" validate:"required,min=5" unique_contraint:"roles_name_key"`
}

func (CreateRoleRequest) TableName() string {
	return "roles"
}

type UpdateRoleRequest struct {
	ID   uuid.UUID `json:"-"`
	Name string    `json:"name" validate:"required"`
}

func (UpdateRoleRequest) TableName() string {
	return "roles"
}
