-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name, api_key) VALUES ($1, $2, $3, encode(sha256(random()::text::bytea), 'hex')) RETURNING *;

-- name: ListAllUsers :many
SELECT * FROM users ORDER BY $1 OFFSET $2 LIMIT $3;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;