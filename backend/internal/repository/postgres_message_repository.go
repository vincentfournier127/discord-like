package repository

import (
	"context"
	"discord-like/internal/api/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresMessageRepository struct {
	db *pgxpool.Pool
}

func NewMessageRepository(db *pgxpool.Pool) *PostgresMessageRepository {
	return &PostgresMessageRepository{
		db: db,
	}
}

func (r *PostgresMessageRepository) Create(ctx context.Context, message *model.Message) (*model.Message, error) {
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO messages (chatroom_id, sender_id, content)
		VALUES($1, $2, $3)
		RETURNING id, created_at`,
		message.Chatroom_ID,
		message.Sender_ID,
		message.Content,
	).Scan(
		&message.ID,
		&message.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (r *PostgresChatroomRepository) GetByChatroomID(ctx context.Context, chatroomID int, limit int, offset int) ([]model.Message, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT *
		FROM messages
		WHERE chatroom_id = $1
		ORDER BY created_at DESC
		LIMIT $2
		OFFSET $3;`,
		chatroomID,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []model.Message

	for rows.Next() {
		var msg model.Message

		err := rows.Scan(
			&msg.ID,
			&msg.Chatroom_ID,
			&msg.Sender_ID,
			&msg.Content,
			&msg.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, msg)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
