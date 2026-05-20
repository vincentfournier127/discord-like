package repository

import (
	"context"
	"discord-like/internal/api/model"
)

type ChatroomRepository interface {
	Create(ctx context.Context, room *model.Chatroom) (*model.Chatroom, error)
	GetByID(ctx context.Context, id int) (model.Chatroom, error)
	GetByUser(ctx context.Context, userID int) ([]model.Chatroom, error)
}
