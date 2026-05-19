package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DATABASE_URL = "postgres://vincent:secretpassword@localhost:5432/discord_like"

func NewPostgresPool() (*pgxpool.Pool, error) {
	ctx := context.Background()

	db, err := pgxpool.New(ctx, DATABASE_URL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
