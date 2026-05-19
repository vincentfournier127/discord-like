package repository

import "github.com/jackc/pgx/v5/pgxpool"

type MessageRepository struct {
	db *pgxpool.Pool
}

func NewMessageRepository(db *pgxpool.Pool) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}
