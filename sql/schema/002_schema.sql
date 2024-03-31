-- +goose Up
ALTER TABLE reviews
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- +goose Down
ALTER TABLE reviews
DROP COLUMN created_at;

