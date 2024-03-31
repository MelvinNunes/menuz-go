CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    code VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    phone_number_code VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) UNIQUE,
    raw_phone_number VARCHAR(255) UNIQUE,
    app_language VARCHAR(255) NOT NULL DEFAULT 'pt',
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);