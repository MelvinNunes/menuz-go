CREATE TABLE promotions (
    id UUID PRIMARY KEY,                            -- Unique identifier for the promotion
    restaurant_id UUID REFERENCES restaurants(id) ON DELETE CASCADE, -- Foreign key to restaurants table
    title VARCHAR(255) NOT NULL,                    -- Title of the promotion
    description TEXT,                               -- Detailed description of the promotion
    promotion_type VARCHAR(255) NOT NULL,           -- Type of promotion (e.g., "buy X, get Y", "percentage off")
    required_items INT,                    -- Number of items that must be purchased
    free_items INT,                        -- Number of free items received
    price DECIMAL(10, 2),                           -- Price for the promotion (if applicable)
    start_date TIMESTAMP NOT NULL,                  -- Start date of the promotion
    end_date TIMESTAMP NOT NULL,                    -- End date of the promotion
    day_of_week VARCHAR(20),                        -- Day of the week for the promotion
    is_active BOOLEAN DEFAULT TRUE,                 -- Status of the promotion (active/inactive)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Updated timestamp
    deleted_at TIMESTAMP                            -- Deleted timestamp (soft delete)
);

CREATE INDEX idx_promotions_deleted_at ON promotions(deleted_at);
