CREATE TABLE menu_items (
    id BIGSERIAL PRIMARY KEY,                     -- ID with auto increment
    menu_id UUID REFERENCES menus(id) ON DELETE CASCADE, -- Foreign key to menus table
    name VARCHAR(255) NOT NULL,                   -- Menu item name
    description TEXT,                             -- Description (optional)
    price DECIMAL(10, 2),                         -- Price of the menu item
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Created timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Updated timestamp
    deleted_at TIMESTAMP                          -- Deleted timestamp (soft delete)
);

CREATE INDEX idx_menu_items_deleted_at ON menu_items(deleted_at);

-- Many-to-many relationship with Category
CREATE TABLE menu_item_categories (
    menu_item_id BIGINT REFERENCES menu_items(id) ON DELETE CASCADE, -- Foreign key to menu_items table
    category_id BIGINT REFERENCES categories(id) ON DELETE CASCADE,  -- Foreign key to categories table
    PRIMARY KEY (menu_item_id, category_id) -- Composite primary key
);

-- Many-to-many relationship with DietaryInfo
CREATE TABLE menu_item_dietary_infos (
    menu_item_id BIGINT REFERENCES menu_items(id) ON DELETE CASCADE, -- Foreign key to menu_items table
    dietary_info_id BIGINT REFERENCES dietary_infos(id) ON DELETE CASCADE, -- Foreign key to dietary_infos table
    PRIMARY KEY (menu_item_id, dietary_info_id) -- Composite primary key
);
