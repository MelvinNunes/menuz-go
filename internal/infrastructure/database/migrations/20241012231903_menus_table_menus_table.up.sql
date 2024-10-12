CREATE TABLE menus (
    id UUID PRIMARY KEY,                           -- UUID as the primary key
    restaurant_id UUID REFERENCES restaurants(id) ON DELETE CASCADE, -- Foreign key to restaurants table
    name VARCHAR(255) NOT NULL,                    -- Menu name (e.g., Lunch, Dinner)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Updated timestamp
    deleted_at TIMESTAMP                           -- Deleted timestamp (soft delete)
);

CREATE INDEX idx_menus_deleted_at ON menus(deleted_at);
