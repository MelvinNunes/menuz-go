CREATE TABLE preferences (
    id BIGSERIAL PRIMARY KEY,                    -- ID with auto increment
    user_id UUID NOT NULL,                       -- User ID (UUID)
    price_range_down DECIMAL(10, 2),             -- Lower bound of price range
    price_range_up DECIMAL(10, 2),               -- Upper bound of price range
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Updated timestamp
    deleted_at TIMESTAMP                         -- Deleted timestamp (soft delete)
);

CREATE INDEX idx_preferences_deleted_at ON preferences(deleted_at);

-- Many-to-many relationship with CuisineTypes
CREATE TABLE preference_cuisine_types (
    preference_id BIGINT REFERENCES preferences(id) ON DELETE CASCADE, -- Foreign key to preferences table
    cuisine_type_id BIGINT REFERENCES cuisine_types(id) ON DELETE CASCADE, -- Foreign key to cuisine_types table
    PRIMARY KEY (preference_id, cuisine_type_id) -- Composite primary key
);

-- Many-to-many relationship with DietaryInfo
CREATE TABLE preference_dietary_infos (
    preference_id BIGINT REFERENCES preferences(id) ON DELETE CASCADE, -- Foreign key to preferences table
    dietary_info_id BIGINT REFERENCES dietary_infos(id) ON DELETE CASCADE, -- Foreign key to dietary_infos table
    PRIMARY KEY (preference_id, dietary_info_id) -- Composite primary key
);