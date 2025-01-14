-- +goose Up
CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   username TEXT NOT NULL,
   email TEXT UNIQUE NOT NULL
);

INSERT INTO users (username, email) VALUES
    ('user1', 'user1@example.com'),
    ('user2', 'user2@example.com');
    ('user3', 'user3@example.com'),
    ('user4', 'user4@example.com');

-- +goose Down
DROP TABLE users;