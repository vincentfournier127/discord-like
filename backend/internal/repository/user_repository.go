package repository

import (
	"context"
	"discord-like/internal/api/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByID(ctx context.Context, id int) (model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
}
