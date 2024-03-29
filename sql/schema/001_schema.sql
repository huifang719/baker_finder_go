-- +goose Up
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_name TEXT NOT NULL UNIQUE,
  user_type TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password_digest TEXT NOT NULL
);

CREATE TABLE bakers(
    id UUID PRIMARY KEY,
    img TEXT NOT NULL,
    name TEXT UNIQUE NOT NULL,
    address TEXT NOT NULL,
    suburb TEXT NOT NULL,
    postcode TEXT NOT NULL,
    contact TEXT NOT NULL,
    specialty TEXT NOT NULL,
    creator UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE reviews(
  id UUID PRIMARY KEY,
  baker_id UUID NOT NULL REFERENCES bakers(id) ON DELETE CASCADE,
  review TEXT NOT NULL,
  rating TEXT NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE bakers;
DROP TABLE reviews;