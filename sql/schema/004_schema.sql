-- +goose Up
DROP TABLE reviews;
DROP TABLE bakers;

CREATE TABLE bakers(
    id UUID PRIMARY KEY,
    img TEXT,
    name TEXT UNIQUE NOT NULL,
    address TEXT,
    suburb TEXT,
    postcode TEXT,
    contact TEXT,
    specialty TEXT,
    creator UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE reviews(
  id UUID PRIMARY KEY,
  baker_id UUID NOT NULL REFERENCES bakers(id) ON DELETE CASCADE,
  review TEXT,
  rating TEXT,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
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