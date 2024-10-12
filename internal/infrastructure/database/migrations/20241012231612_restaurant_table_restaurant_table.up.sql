CREATE TABLE restaurants (
    id UUID PRIMARY KEY,                           -- UUID as the primary key
    name VARCHAR(255) NOT NULL,                    -- Restaurant name
    address TEXT NOT NULL,                         -- Address
    price_range_down DECIMAL(10, 2),               -- Lower bound of price range
    price_range_up DECIMAL(10, 2),                 -- Upper bound of price range
    latitude DECIMAL(10, 8),                       -- Latitude for geolocation
    longitude DECIMAL(11, 8),                      -- Longitude for geolocation
    is_approved BOOLEAN DEFAULT FALSE,             -- Approval status
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Updated timestamp
    deleted_at TIMESTAMP                           -- Deleted timestamp (soft delete)
);

CREATE INDEX idx_restaurants_deleted_at ON restaurants(deleted_at);

-- Many-to-many relationship with CuisineTypes
CREATE TABLE restaurant_cuisine_types (
    restaurant_id UUID REFERENCES restaurants(id) ON DELETE CASCADE, -- Foreign key to restaurants table
    cuisine_type_id BIGINT REFERENCES cuisine_types(id) ON DELETE CASCADE, -- Foreign key to cuisine_types table
    PRIMARY KEY (restaurant_id, cuisine_type_id) -- Composite primary key
);