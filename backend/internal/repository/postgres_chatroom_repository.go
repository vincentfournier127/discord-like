package repository

import (
	"context"
	"discord-like/internal/api/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresChatroomRepository struct {
	db *pgxpool.Pool
}

func NewChatroomRepository(db *pgxpool.Pool) *PostgresChatroomRepository {
	return &PostgresChatroomRepository{
		db: db,
	}
}

func (r *PostgresChatroomRepository) Create(ctx context.Context, c *model.Chatroom) (*model.Chatroom, error) {
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

func (r *PostgresChatroomRepository) GetByID(ctx context.Context, chatroomID int) (model.Chatroom, error) {
	var c model.Chatroom

	err := r.db.QueryRow(
		ctx,
		`SELECT id, name, created_at
		FROM chatroom
		WHERE id = $1`,
		chatroomID,
	).Scan(
		&c.ID,
		&c.Name,
		&c.CreatedAt,
	)

	if err != nil {
		return c, err
	}

	return c, nil
}

func (r *PostgresChatroomRepository) GetByUserID(ctx context.Context, userID int) ([]model.Chatroom, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT c.*
		FROM chatrooms c
		JOIN memberships m ON m.chatroom_id = c.id
		WHERE m.user_id = $1;`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var chatrooms []model.Chatroom

	for rows.Next() {
		var c model.Chatroom

		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		chatrooms = append(chatrooms, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chatrooms, nil
}
