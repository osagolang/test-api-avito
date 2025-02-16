
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    coins INT NOT NULL DEFAULT 0
    );

-- +migrate Down
DROP TABLE IF EXISTS users;