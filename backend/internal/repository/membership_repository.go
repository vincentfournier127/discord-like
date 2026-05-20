package repository

import (
	"context"
	"discord-like/internal/api/model"
)

type MembershipRepository interface {
	AddUserToChatroom(ctx context.Context, userID int, chatroomID int) error
	RemoveUserFromChatroom(ctx context.Context, userID int, chatroomID int) error
	GetUsersByChatroomID(ctx context.Context, chatroomID int) ([]model.User, error)
}
