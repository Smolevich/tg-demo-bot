-- +goose Up
CREATE TABLE IF NOT EXISTS messages (
    id BIGINT GENERATED ALWAYS AS IDENTITY,
    chat_id BIGINT,
    username VARCHAR(200),
    user_id BIGINT,
    text TEXT,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE messages;
