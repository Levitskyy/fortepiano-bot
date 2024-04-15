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

func (s *SubscriptionPostgreStorage) UpdateEndDate(ctx context.Context, subscription model.Subscription, days int) error {
	conn, err := s.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	endDate := time.Now().Add(time.Hour * 24 * time.Duration(days))
	err = s.db.Get(&subscription, "SELECT * FROM subscriptions WHERE user_id=$1 AND group_id=$2", subscription.UserId, subscription.GroupId)
	if err == nil {
		var resEndDate time.Time
		subEndDate := subscription.EndDate.Add(time.Hour * 24 * time.Duration(days))

		if subEndDate.After(endDate) {
			resEndDate = subEndDate
		} else {
			resEndDate = endDate
		}
		if _, err := conn.ExecContext(
			ctx,
			`UPDATE subscriptions SET end_date = $1 WHERE id = $2;`,
			resEndDate,
			subscription.Id,
		); err != nil {
			return err
		}
	} else {
		if _, err := conn.ExecContext(
			ctx,
			`INSERT INTO subscriptions (user_id, group_id, end_date) VALUES ($1, $2, $3);`,
			subscription.UserId,
			subscription.GroupId,
			endDate,
		); err != nil {
			return err
		}
	}

	return nil
}

func (s *SubscriptionPostgreStorage) IsSubActive(ctx context.Context, subscription model.Subscription) (bool, error) {
	err := s.db.Get(&subscription, "SELECT * FROM subscriptions WHERE user_id=$1 AND group_id=$2", subscription.UserId, subscription.GroupId)
	if err != nil {
		return false, nil
	}
	if time.Now().After(subscription.EndDate) {
		return false, nil
	}
	return true, nil
}

func (s *SubscriptionPostgreStorage) GetUserSubs(ctx context.Context, subscription model.Subscription) ([]model.Subscription, error) {
	subs := []model.Subscription{}

	err := s.db.Select(&subs, "SELECT * FROM subscriptions WHERE user_id=$1 AND end_date > $2::timestamp", subscription.UserId, time.Now().Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	return subs, nil

}
