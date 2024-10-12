CREATE TABLE dietary_infos (
    id BIGSERIAL PRIMARY KEY,                    -- ID with auto increment
    name VARCHAR(255) NOT NULL UNIQUE,           -- Unique name
    description TEXT,                            -- Description (optional)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Updated timestamp
    deleted_at TIMESTAMP
);

CREATE INDEX idx_dietary_infos_deleted_at ON dietary_infos(deleted_at);