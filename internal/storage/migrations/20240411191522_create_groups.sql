-- +goose Up
-- +goose StatementBegin
CREATE TABLE groups (
    id          BIGINT       PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    invite_link VARCHAR(255) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS groups;
-- +goose StatementEnd
