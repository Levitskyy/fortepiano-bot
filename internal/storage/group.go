package storage

import (
	"context"
	"fmt"

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

func (s *GroupPostgreStorage) GetId(ctx context.Context, group model.Group) (int64, error) {
	err := s.db.Get(&group, "SELECT * FROM groups WHERE name=$1", group.Name)
	if err == nil {
		return group.Id, nil
	}

	return 0, fmt.Errorf("no group with name %s", group.Name)
}

func (s *GroupPostgreStorage) GetLink(ctx context.Context, group model.Group) (string, error) {
	err := s.db.Get(&group, "SELECT * FROM groups WHERE id=$1", group.Id)
	if err == nil {
		return group.InviteLink, nil
	}

	return "", fmt.Errorf("no group with id %v", group.Id)
}

func (s *GroupPostgreStorage) GetName(ctx context.Context, group model.Group) (string, error) {
	err := s.db.Get(&group, "SELECT * FROM groups WHERE id=$1", group.Id)
	if err == nil {
		return group.Name, nil
	}

	return "", fmt.Errorf("no group with id %v", group.Id)
}
