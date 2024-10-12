CREATE TABLE item_promotions (
    id UUID PRIMARY KEY,                            -- Unique identifier for the item promotion
    menu_item_id BIGINT NOT NULL REFERENCES menu_items(id) ON DELETE CASCADE, -- Foreign key to menu_items table
    title VARCHAR(255) NOT NULL,                    -- Title of the promotion
    description TEXT,                               -- Detailed description of the promotion (optional)
    discount_amount DECIMAL(10, 2) NOT NULL,       -- Discount amount (e.g., in currency)
    discount_type VARCHAR(50) NOT NULL,             -- Type of discount (e.g., "percentage", "fixed")
    start_date TIMESTAMP NOT NULL,                  -- Start date of the promotion
    end_date TIMESTAMP NOT NULL,                    -- End date of the promotion
    is_active BOOLEAN DEFAULT TRUE,                 -- Status of the promotion (active/inactive)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Updated timestamp
    deleted_at TIMESTAMP                            -- Deleted timestamp (soft delete)
);

CREATE INDEX idx_item_promotions_deleted_at ON item_promotions(deleted_at);
