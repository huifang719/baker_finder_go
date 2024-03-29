-- +goose Up
ALTER TABLE reviews DROP COLUMN rating;
ALTER TABLE reviews ADD COLUMN rating TEXT;

-- +goose Down
ALTER TABLE reviews DROP COLUMN rating;
ALTER TABLE reviews ADD COLUMN rating NUMERIC;