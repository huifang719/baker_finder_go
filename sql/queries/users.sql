-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, user_name, user_type, email, password_digest)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: CreateBaker :one
INSERT INTO bakers (id, img, name, address, suburb, postcode, contact, specialty, creator)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: CreateReview :one
INSERT INTO reviews (id, baker_id, review, rating, user_name)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteBaker :one
DELETE FROM bakers
WHERE id = $1
RETURNING *;

-- name: DeleteReview :one
DELETE FROM reviews
WHERE id = $1
RETURNING *;

-- name: GetBakersByPostcode :many
SELECT * FROM bakers
WHERE postcode = $1;

-- name: GetAllReviews :many
SELECT * FROM reviews
WHERE baker_id = $1;

-- name: UpdateBaker :one
UPDATE bakers SET img = $2, name = $3, address = $4, suburb = $5, postcode = $6, contact = $7, specialty = $8, creator = $9
WHERE id = $1
RETURNING *;

-- name: GetBakerById :one
SELECT * FROM bakers
WHERE id = $1;
