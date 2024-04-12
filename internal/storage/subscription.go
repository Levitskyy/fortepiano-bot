package storage

import (
	"context"
	"time"

	"github.com/Levitskyy/fortepiano-bot/internal/model"
	"github.com/jmoiron/sqlx"
)

type SubscriptionPostgreStorage struct {
	db *sqlx.DB
}

func NewSubscriptionStorage(db *sqlx.DB) *SubscriptionPostgreStorage {
	return &SubscriptionPostgreStorage{db: db}
}

func (s *SubscriptionPostgreStorage) Add(ctx context.Context, subscription model.Subscription) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := s.db.ExecContext(
		ctx,
		`INSERT INTO subscriptions (user_id, group_id, end_date)
			VALUES ($1, $2, $3)`,
		subscription.UserId,
		subscription.GroupId,
		subscription.EndDate,
	); err != nil {
		return err
	}

	return nil
}

func (s *SubscriptionPostgreStorage) UpdateEndDate(ctx context.Context, subscription model.Subscription, endDate time.Time) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		`UPDATE subscriptions SET end_date = $1 WHERE id = $2;`,
		endDate,
		subscription.Id,
	); err != nil {
		return err
	}

	return nil
}
