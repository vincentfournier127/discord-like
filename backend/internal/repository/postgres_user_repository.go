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
		RETURNING id`,
		u.Name,
	).Scan(
		&u.ID,
	)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int) (model.User, error) {

	var user model.User

	err := r.db.QueryRow(
		ctx,
		`SELECT id, name
		FROM users 
		WHERE id = $1`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *PostgresUserRepository) GetAll(ctx context.Context) ([]model.User, error) {

	rows, err := r.db.Query(
		ctx,
		`SELECT id, name
		FROM users`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var user model.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
