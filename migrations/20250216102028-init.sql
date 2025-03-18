
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    ID SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    coins INT NOT NULL DEFAULT 0
    );

CREATE TABLE IF NOT EXISTS inventory (
    user_id INT NOT NULL REFERENCES users(id),
    type VARCHAR(255) NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    PRIMARY KEY (user_id, type)
    );

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    from_user INT REFERENCES users(id),
    to_user INT REFERENCES users(id),
    amount INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS items (
    type VARCHAR(255) PRIMARY KEY UNIQUE,
    price INT NOT NULL
    );

INSERT INTO items (type, price) VALUES
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500);

-- +migrate Down
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS inventory;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS users;