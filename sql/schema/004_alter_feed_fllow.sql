-- +goose Up
ALTER TABLE feed_follows ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE feed_follows DROP COLUMN last_fetched_at;
