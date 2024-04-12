-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions (
    id       SERIAL PRIMARY KEY,
    user_id  BIGINT REFERENCES users (id),
    group_id BIGINT REFERENCES groups (id),
    end_date TIMESTAMP   NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions;
-- +goose StatementEnd
