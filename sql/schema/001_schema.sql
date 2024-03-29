-- +goose Up
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_name TEXT NOT NULL,
  user_type TEXT NOT NULL,
  email TEXT NOT NULL,
  password_digest TEXT NOT NULL
);

CREATE TABLE bakers(
    id SERIAL PRIMARY KEY,
    img TEXT,
    name TEXT,
    address TEXT,
    suburb TEXT,
    postcode TEXT,
    contact TEXT,
    specialty TEXT,
    creator TEXT
);

CREATE TABLE reviews(
  id SERIAL PRIMARY KEY,
  baker_id NUMERIC,
  review TEXT,
  rating NUMERIC,
  user_name TEXT
);

-- +goose Down

DROP TABLE bakers;
DROP TABLE reviews;