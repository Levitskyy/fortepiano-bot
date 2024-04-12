package storage

import (
	"context"

	"github.com/Levitskyy/fortepiano-bot/internal/model"

	"github.com/jmoiron/sqlx"
)

type GroupPostgreStorage struct {
	db *sqlx.DB
}

func NewGroupStorage(db *sqlx.DB) *GroupPostgreStorage {
	return &GroupPostgreStorage{db: db}
}

func (s *GroupPostgreStorage) Add(ctx context.Context, group model.Group) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := s.db.ExecContext(
		ctx,
		`INSERT INTO groups (id, name, invite_link)
			VALUES ($1, $2, $3)`,
		group.Id,
		group.Name,
		group.InviteLink,
	); err != nil {
		return err
	}

	return nil
}
