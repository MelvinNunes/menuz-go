CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,                    -- ID with auto increment
    name VARCHAR(255) NOT NULL UNIQUE,           -- Unique name
    description TEXT,                            -- Description (optional)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Updated timestamp
    deleted_at TIMESTAMP                       -- Deleted timestamp (soft delete)
);

CREATE INDEX idx_categories_deleted_at ON categories(deleted_at);