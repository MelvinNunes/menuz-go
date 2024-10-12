CREATE TABLE saved_menus (
    id BIGSERIAL PRIMARY KEY,                      -- ID with auto increment
    user_id UUID NOT NULL,                         -- Foreign key for user (UUID)
    menu_id UUID REFERENCES menus(id) ON DELETE CASCADE, -- Foreign key to menus table
    notes TEXT,                                    -- Optional notes for the saved menu
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Updated timestamp
    deleted_at TIMESTAMP                           -- Deleted timestamp (soft delete)
);

CREATE INDEX idx_saved_menus_deleted_at ON saved_menus(deleted_at);
