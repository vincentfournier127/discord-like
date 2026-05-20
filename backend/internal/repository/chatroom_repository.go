package repository

import (
	"context"
	"discord-like/internal/api/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatroomRepository struct {
	db *pgxpool.Pool
}

func NewChatroomRepository(db *pgxpool.Pool) *ChatroomRepository {
	return &ChatroomRepository{
		db: db,
	}
}

func (r *ChatroomRepository) CreateChatroom(ctx context.Context, c *model.Chatroom) (*model.Chatroom, error) {
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO chatroom (name)
		VALUES($1)
		RETURNING id`,
		c.Name,
	).Scan(
		&c.ID,
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}
