-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, user_name, user_type, email, password_digest)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: CreateBaker :one
INSERT INTO bakers (img, name, address, suburb, postcode, contact, specialty, creator)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

