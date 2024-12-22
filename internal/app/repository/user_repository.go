package repository

import (
	"context"

	"github.com/umardev500/pos/internal/app/contract"
	"github.com/umardev500/pos/internal/app/model"
	"github.com/umardev500/pos/pkg"
)

type userRepository struct {
	db *pkg.GormInstance
}

func NewUserRepository(db *pkg.GormInstance) contract.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User

	if err := u.db.GetConn(ctx).Preload("Roles").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
