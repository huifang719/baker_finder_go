-- +goose Up
DROP TABLE bakers;
DROP TABLE reviews;
CREATE TABLE bakers(
    id UUID PRIMARY KEY,
    img TEXT,
    name TEXT UNIQUE NOT NULL,
    address TEXT,
    suburb TEXT,
    postcode TEXT,
    contact TEXT,
    specialty TEXT,
    creator UUID REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE reviews(
  id UUID PRIMARY KEY,
  baker_id UUID REFERENCES bakers(id) ON DELETE CASCADE,
  review TEXT,
  rating TEXT,
  user_id UUID REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE bakers;
DROP TABLE reviews;
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
  rating TEXT,
  user_name TEXT
);