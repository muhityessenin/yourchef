-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    address VARCHAR,
    registration_date DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR NOT NULL CHECK (role IN ('admin', 'user', 'premium_user', 'guest'))
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd