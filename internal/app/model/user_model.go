package model

import (
	"github.com/umardev500/pos/pkg"
)

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password_hash string `json:"password_hash,omitempty"`
	Roles         []Role `json:"roles" gorm:"many2many:user_roles"`
	pkg.TimeCommon
}

type CreateUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
}

func (CreateUserReq) TableName() string {
	return "users"
}

type UserFindParams struct{}
