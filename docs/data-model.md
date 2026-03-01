# Data Model Documentation

This document describes the domain models and key business logic in Forage.

## Items & Units

### Item
Represents a product that can be inventoried, purchased, or used in meals. Each item can belong to multiple categories through `ItemCategoryMapping`.

- **id**: Unique identifier
- **name**: Product name (must be unique)

### Unit
Represents a unit of measurement (e.g., "lbs", "cups", "slices"). Units are optional and referenced by inventory items and store items.

- **id**: Unique identifier
- **name**: Unit name (must be unique)

### ItemCategory
Categories for classification (e.g., foods, produce, dairy). Items can have multiple categories.

- **id**: Unique identifier
- **name**: Category name (must be unique)

### ItemCategoryMapping
Many-to-many junction table linking items to categories.

## Inventory

### InventoryItem
Tracks the current state of items purchased and available for use. Each item has at most one inventory record (1:1 relationship with Item).

- **id**: Unique identifier
- **item_id**: Reference to the item
- **unit_id**: Optional unit of measurement for tracking quantities
- **stored_amount**: Current quantity in inventory (must be ≥ 0)
  - When null: the item is presense-based and its amount is not tracked

## Meals & Ingredients

### Meal
A recipe that serves a specific number of people/portions.

- **id**: Unique identifier
- **name**: Meal/recipe name
- **meal_category_id**: Reference to meal category (e.g., breakfast, lunch, dinner)
- **servings**: Number of servings this recipe makes (must be > 0)

### MealCategory
Categories for organizing meals (e.g., breakfast, lunch, dinner, snacks). 

- **id**: Unique identifier
- **name**: Category name (must be unique)

### MealIngredient
Links items to meals with quantity information. Represents the ingredients required for a recipe.

- **id**: Unique identifier
- **meal_id**: Reference to the meal
- **ingredient_id**: Reference to the item being used as an ingredient
- **unit_id**: Optional unit of measurement for the required amount
- **required_amount**: Quantity of ingredient required in the recipe (must be > 0 if specified)
  - When null: ingredient is "to taste", a topping/condiment, on the side, or otherwise doesn't require precise quantity tracking
- **is_seasoning**: Boolean flag; true if the ingredient can logically be grouped as a seasoning
  - Used for filtering/grouping in UI
- **depletes_slowly**: Boolean flag; true if the ingredient is typically used in small amounts and depletes over time
  - Examples: spices, condiments
- **alternate_ingredient_id**: Optional reference to another item
  - Represents ingredient substitution support (e.g., butter ↔ oil, salt ↔ sea salt)

## Stores & Store Items

### Store
A retail location that sells items.

- **id**: Unique identifier
- **name**: Store name (must be unique)

### StoreItem
Links items to stores with purchase and inventory unit information. Represents products sold at a particular store.

- **id**: Unique identifier
- **store_id**: Reference to the store
- **item_id**: Reference to the item
- **purchase_unit_id**: Optional unit by which the item is sold (e.g., "loaf" for bread)
- **inventory_unit_id**: Optional unit used for internal inventory tracking (e.g., "slice" for a loaf of bread)
- **inventory_units_per_purchase**: Conversion factor from purchase unit to inventory unit (e.g., 16 for a loaf of bread)
  - Default: 1.0 (purchase and inventory units are the same)
  - Must be > 0
- **purchased_by_decimal**: Boolean; true if item can be purchased in fractional amounts
  - Examples: true for items sold by weight
  - Constraint: If true, then inventory_units_per_purchase must equal 1
- **is_favorite**: Boolean; indicates if this item is favorited by the user
- **store_traversal_order**: Integer representing the order items appear in while traversing the store layout

## Shopping Lists

### ShoppingList
A collection of items to purchase.

- **id**: Unique identifier
- **name**: User-defined name for the shopping list
- **created_at**: Timestamp when the list was created (defaults to CURRENT_TIMESTAMP)

### ShoppingListItem
Individual items added to a shopping list. Each item is either a store item or a custom user-defined item.

- **id**: Unique identifier
- **shopping_list_id**: Reference to the shopping list
- **store_item_id**: Optional reference to a store item catalog entry
  - When null: use custom_item_name instead
- **purchase_quantity**: Quantity to purchase
  - Optional; null is allowed (e.g., for items like "Bread" where quantity is implicit)
  - Must be > 0 if specified
- **note**: Optional user annotation for the item
  - Examples: "Green bananas", "Organic preferred"
- **custom_item_name**: Optional custom item name for shopping list items that don't correspond to a catalog store item
  - Constraint: Either store_item_id OR custom_item_name must be non-null (XOR logic)
  - Used for ad-hoc items not in the system
- **shopping_list_order**: Integer representing the order items appear in the shopping list

## Item Quantity Representation Patterns
Items are represented differently based on their quantity/unit combination:
- Measured: has quantity + unit ("6 slices of bread", "1.25 lbs of deli turkey")
- Counted: has quantity, no unit ("3 chicken breasts", "4 bananas")  
- Presence-based: no quantity or unit ("BBQ sauce", "Garlic powder")
