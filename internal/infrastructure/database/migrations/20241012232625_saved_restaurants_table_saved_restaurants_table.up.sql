CREATE TABLE saved_restaurants (
    id BIGSERIAL PRIMARY KEY,                      -- ID with auto increment
    user_id UUID NOT NULL,                         -- Foreign key for user (UUID)
    restaurant_id UUID REFERENCES restaurants(id) ON DELETE CASCADE, -- Foreign key to restaurants table
    notes TEXT,                                    -- Optional notes for the saved restaurant
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Updated timestamp
    deleted_at TIMESTAMP                           -- Deleted timestamp (soft delete)
);

CREATE INDEX idx_saved_restaurants_deleted_at ON saved_restaurants(deleted_at);
