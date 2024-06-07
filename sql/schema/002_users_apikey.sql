-- +goose Up
ALTER TABLE users ADD COLUMN IF NOT EXISTS api_key VARCHAR(64) DEFAULT encode(sha256(random()::text::bytea), 'hex'); 

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;