package storage

import (
	"context"

	"github.com/Levitskyy/fortepiano-bot/internal/model"

	"github.com/jmoiron/sqlx"
)

type UserPostgreStorage struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *UserPostgreStorage {
	return &UserPostgreStorage{db: db}
}

func (s *UserPostgreStorage) Add(ctx context.Context, user model.User) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		`INSERT INTO users (id, name)
			VALUES ($1, $2);`,
		user.Id,
		user.Name,
	); err != nil {
		return err
	}

	return nil
}

func (s *UserPostgreStorage) UpdateEmail(ctx context.Context, user model.User, email string) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		`UPDATE users SET email = $1 WHERE id = $2
			ON CONFLICT DO NOTHING;`,
		email,
		user.Id,
	); err != nil {
		return err
	}

	return nil
}
