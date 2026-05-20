package repository

import (
	"context"
	"discord-like/internal/api/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}

}

func (r *PostgresUserRepository) Create(ctx context.Context, u *model.User) (*model.User, error) {

	err := r.db.QueryRow(
		ctx,
		`INSERT INTO users (name) 
		VALUES($1) 
		RETURNING id, created_at`,
		u.Name,
	).Scan(
		&u.ID,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int) (model.User, error) {

	var u model.User

	err := r.db.QueryRow(
		ctx,
		`SELECT id, name, created_at
		FROM users 
		WHERE id = $1`,
		id,
	).Scan(
		&u.ID,
		&u.Name,
		&u.CreatedAt,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}

func (r *PostgresUserRepository) GetAll(ctx context.Context) ([]model.User, error) {

	rows, err := r.db.Query(
		ctx,
		`SELECT id, name, created_at
		FROM users`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User

		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
