CREATE TABLE reviews (
    id UUID PRIMARY KEY,                           -- UUID as the primary key
    user_id UUID NOT NULL,                         -- Foreign key for user (UUID)
    restaurant_id UUID NOT NULL,                   -- Foreign key for restaurant (UUID)
    menu_item_id BIGINT NOT NULL REFERENCES menu_items(id) ON DELETE CASCADE, -- Foreign key to menu_items table
    rating INT, -- Rating constraint (1 to 5)
    comment TEXT,                                  -- Optional comment
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Updated timestamp
    deleted_at TIMESTAMP                           -- Deleted timestamp (soft delete)
);

CREATE INDEX idx_reviews_deleted_at ON reviews(deleted_at);
