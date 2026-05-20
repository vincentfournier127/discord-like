package repository

import (
	"context"
	"discord-like/internal/api/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresMembershipRepository struct {
	db *pgxpool.Pool
}

func NewPostgresMembershipRepostory(db *pgxpool.Pool) *PostgresMembershipRepository {
	return &PostgresMembershipRepository{
		db: db,
	}
}

func (r *PostgresMembershipRepository) AddUserToChatroom(ctx context.Context, userID int, chatroomID int) error {
	_, err := r.db.Exec(
		ctx,
		`INSERT INTO memberships (user_id, chatroom_id)
	 	VALUES ($1,$2)`,
		userID,
		chatroomID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresMembershipRepository) RemoveUserFromChatroom(ctx context.Context, userID int, chatroomID int) error {
	_, err := r.db.Exec(
		ctx,
		`DELETE FROM memberships
		WHERE user_id = $1
		AND chatroomID = $2`,
		userID,
		chatroomID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresMembershipRepository) GetUserByChatroomID(ctx context.Context, chatroomID int) ([]model.User, error) {
	rows, err := r.db.Query(
		ctx,
		`SELECT u.*
		FROM users u
		JOIN memberships m ON m.user_id = u.id
		WHERE m.chatroom_id = $1;`,
		chatroomID,
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
