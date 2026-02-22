CREATE TABLE units (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE item_categories (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE meal_categories (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE items (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Many-to-many junction table: items can belong to multiple categories
CREATE TABLE item_category_mappings (
    item_id INT NOT NULL,
    category_id INT NOT NULL,
    CONSTRAINT item_category_mappings_fk_item FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE,
    CONSTRAINT item_category_mappings_fk_category FOREIGN KEY (category_id) REFERENCES item_categories(id) ON DELETE RESTRICT,
    PRIMARY KEY (item_id, category_id)
);

-- 1:1 relationship: each item has at most one inventory record
CREATE TABLE inventory_items (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    item_id INT NOT NULL UNIQUE,
    unit_id INT,
    stored_amount NUMERIC(10, 4),
    CONSTRAINT inventory_items_fk_item FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE RESTRICT,
    CONSTRAINT inventory_items_fk_unit FOREIGN KEY (unit_id) REFERENCES units(id) ON DELETE RESTRICT,
    CONSTRAINT inventory_items_stored_amount_check CHECK (stored_amount >= 0)
);

CREATE TABLE stores (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE store_items (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    item_id INT NOT NULL,
    store_id INT NOT NULL,
    purchase_unit_id INT,
    inventory_unit_id INT,
    inventory_units_per_purchase NUMERIC(10, 4) NOT NULL DEFAULT 1.0,
    purchased_by_decimal BOOLEAN NOT NULL DEFAULT FALSE,
    store_traversal_order INT NOT NULL,
    CONSTRAINT store_items_fk_item FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE RESTRICT,
    CONSTRAINT store_items_fk_store FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE,
    CONSTRAINT store_items_fk_purchase_unit FOREIGN KEY (purchase_unit_id) REFERENCES units(id) ON DELETE RESTRICT,
    CONSTRAINT store_items_fk_inventory_unit FOREIGN KEY (inventory_unit_id) REFERENCES units(id) ON DELETE RESTRICT,
    CONSTRAINT store_items_inventory_units_per_purchase_check CHECK (inventory_units_per_purchase > 0),
    CONSTRAINT store_items_purchased_by_decimal_check CHECK (purchased_by_decimal = false OR inventory_units_per_purchase = 1)
);

CREATE TABLE meals (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    meal_category_id INT NOT NULL,
    servings INT NOT NULL,
    CONSTRAINT meals_fk_meal_category FOREIGN KEY (meal_category_id) REFERENCES meal_categories(id) ON DELETE RESTRICT,
    CONSTRAINT meals_servings_check CHECK (servings > 0)
);

CREATE TABLE meal_ingredients (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    meal_id INT NOT NULL,
    ingredient_id INT NOT NULL,
    unit_id INT,
    required_amount NUMERIC(10, 4),
    is_seasoning BOOLEAN NOT NULL DEFAULT FALSE,
    alternate_ingredient_id INT,
    CONSTRAINT meal_ingredients_fk_meal FOREIGN KEY (meal_id) REFERENCES meals(id) ON DELETE CASCADE,
    CONSTRAINT meal_ingredients_fk_ingredient FOREIGN KEY (ingredient_id) REFERENCES items(id) ON DELETE RESTRICT,
    CONSTRAINT meal_ingredients_fk_unit FOREIGN KEY (unit_id) REFERENCES units(id) ON DELETE RESTRICT,
    CONSTRAINT meal_ingredients_fk_alternate_ingredient FOREIGN KEY (alternate_ingredient_id) REFERENCES items(id) ON DELETE RESTRICT,
    CONSTRAINT meal_ingredients_required_amount_check CHECK (required_amount > 0)
);

CREATE TABLE shopping_lists (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- store_item_id is optional: if null, use custom_item_name for manual items
CREATE TABLE shopping_list_items (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    shopping_list_id INT NOT NULL,
    store_item_id INT,
    purchase_quantity NUMERIC(10, 4),
    note TEXT,
    custom_item_name TEXT,
    shopping_list_order INT NOT NULL,
    CONSTRAINT shopping_list_items_fk_shopping_list FOREIGN KEY (shopping_list_id) REFERENCES shopping_lists(id) ON DELETE CASCADE,
    CONSTRAINT shopping_list_items_fk_store_item FOREIGN KEY (store_item_id) REFERENCES store_items(id) ON DELETE RESTRICT,
    CONSTRAINT shopping_list_items_item_source_check CHECK (
        (store_item_id IS NOT NULL AND custom_item_name IS NULL) OR 
        (store_item_id IS NULL AND custom_item_name IS NOT NULL)
    ),
    CONSTRAINT shopping_list_items_purchase_quantity_check CHECK (purchase_quantity > 0)
);

CREATE INDEX idx_item_category_mappings_category_id ON item_category_mappings(category_id);
CREATE INDEX idx_inventory_items_unit_id ON inventory_items(unit_id);
CREATE INDEX idx_store_items_item_id ON store_items(item_id);
CREATE INDEX idx_store_items_store_id ON store_items(store_id);
CREATE INDEX idx_store_items_purchase_unit_id ON store_items(purchase_unit_id);
CREATE INDEX idx_store_items_inventory_unit_id ON store_items(inventory_unit_id);
CREATE INDEX idx_meals_meal_category_id ON meals(meal_category_id);
CREATE INDEX idx_meal_ingredients_meal_id ON meal_ingredients(meal_id);
CREATE INDEX idx_meal_ingredients_ingredient_id ON meal_ingredients(ingredient_id);
CREATE INDEX idx_meal_ingredients_unit_id ON meal_ingredients(unit_id);
CREATE INDEX idx_meal_ingredients_alternate_ingredient_id ON meal_ingredients(alternate_ingredient_id);
CREATE INDEX idx_shopping_list_items_shopping_list_id ON shopping_list_items(shopping_list_id);
CREATE INDEX idx_shopping_list_items_store_item_id ON shopping_list_items(store_item_id);
CREATE INDEX idx_shopping_list_items_order ON shopping_list_items(shopping_list_order);
CREATE INDEX idx_items_name ON items(name);
CREATE INDEX idx_meals_name ON meals(name);
CREATE INDEX idx_stores_name ON stores(name);
CREATE INDEX idx_units_name ON units(name);
