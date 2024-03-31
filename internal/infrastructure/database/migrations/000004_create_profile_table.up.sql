CREATE TABLE IF NOT EXISTS profiles (
    id              SERIAL PRIMARY KEY,
    user_id         UUID NOT NULL,
    first_name      VARCHAR(255),
    last_name       VARCHAR(255),
    gender          VARCHAR(255),
    avatar          VARCHAR(255),
    date_of_birth   DATE,
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP,
    deleted_at      TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);