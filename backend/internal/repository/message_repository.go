package repository

import (
	"context"
	"discord-like/internal/api/model"
)

type MessageRepository interface {
	Create(ctx context.Context, msg *model.Message) (*model.Message, error)
	GetByChatroomID(ctx context.Context, chatroomID int, limit int, offset int) ([]model.Message, error)
}
