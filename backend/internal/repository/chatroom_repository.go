package repository

import "github.com/jackc/pgx/v5/pgxpool"

type ChatroomRepository struct {
	db *pgxpool.Pool
}

func NewChatroomRepository(db *pgxpool.Pool) *ChatroomRepository {
	return &ChatroomRepository{
		db: db,
	}
}
