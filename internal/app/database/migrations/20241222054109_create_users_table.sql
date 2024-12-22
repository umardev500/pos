-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  email VARCHAR(60) NOT NULL UNIQUE,
  username VARCHAR(50) NOT NULL UNIQUE,
  password_hash VARCHAR(60) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NULL DEFAULT NULL,
  deleted_at TIMESTAMPTZ NULL DEFAULT NULL
);


-- Seed data
INSERT INTO users (id, email, username, password_hash) VALUES
  ('00000000-0000-0000-0000-000000000001', 'admin@email.com', 'umardev500', 'hashed_pass'),
  ('00000000-0000-0000-0000-000000000002', 'user@email.com', 'username2', 'hashed_pass');


-- +goose Down
DROP TABLE IF EXISTS users;
