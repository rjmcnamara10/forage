package models

import "time"

// Item and Unit
// Item represents a product that can be inventoried, purchased, or used in meals
type Item struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Unit struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type ItemCategory struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type ItemCategoryMapping struct {
	ItemID     int `json:"item_id" db:"item_id"`
	CategoryID int `json:"category_id" db:"category_id"`
}

// InventoryItem tracks current state of items purchased and available for use
type InventoryItem struct {
	ID           int      `json:"id" db:"id"`
	ItemID       int      `json:"item_id" db:"item_id"`
	UnitID       *int     `json:"unit_id" db:"unit_id"`
	StoredAmount *float64 `json:"stored_amount" db:"stored_amount"` // Current quantity in inventory
}

// Meal
type Meal struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	MealCategoryID int    `json:"meal_category_id" db:"meal_category_id"`
	Servings       int    `json:"servings" db:"servings"`
}

type MealIngredient struct {
	ID                    int      `json:"id" db:"id"`
	MealID                int      `json:"meal_id" db:"meal_id"`
	IngredientID          int      `json:"ingredient_id" db:"ingredient_id"`
	UnitID                *int     `json:"unit_id" db:"unit_id"`
	RequiredAmount        *float64 `json:"required_amount" db:"required_amount"`                 // Quantity of the ingredient used in the meal or called for in the recipe. When null, the ingredient is to taste, a topping, on the side, or otherwise doesn't require precise quantity tracking
	IsSeasoning           bool     `json:"is_seasoning" db:"is_seasoning"`                       // True if the ingredient can be logically grouped as a seasoning
	DepletesSlowly        bool     `json:"depletes_slowly" db:"depletes_slowly"`                 // True if the ingredient is typically used in small amounts and depletes slowly (e.g., spices, condiments)
	AlternateIngredientID *int     `json:"alternate_ingredient_id" db:"alternate_ingredient_id"` // Optional, used when an ingredient can be substituted with another
}

type MealCategory struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Store
type StoreItem struct {
	ID                        int     `json:"id" db:"id"`
	StoreID                   int     `json:"store_id" db:"store_id"`
	ItemID                    int     `json:"item_id" db:"item_id"`
	PurchaseUnitID            *int    `json:"purchase_unit_id" db:"purchase_unit_id"`                         // Unit by which the item is sold
	InventoryUnitID           *int    `json:"inventory_unit_id" db:"inventory_unit_id"`                       // Unit used for inventory tracking
	InventoryUnitsPerPurchase float64 `json:"inventory_units_per_purchase" db:"inventory_units_per_purchase"` // Number of inventory units per purchase unit
	PurchasedByDecimal        bool    `json:"purchased_by_decimal" db:"purchased_by_decimal"`                 // True if item can be purchased in fractional amounts (e.g., weight)
	IsFavorite                bool    `json:"is_favorite" db:"is_favorite"`                                   // True if item is favorited by the user
	StoreTraversalOrder       int     `json:"store_traversal_order" db:"store_traversal_order"`               // Order that items are encountered while traversing store
}

type Store struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Shopping List
type ShoppingListItem struct {
	ID                int      `json:"id" db:"id"`
	ShoppingListID    int      `json:"shopping_list_id" db:"shopping_list_id"`
	StoreItemID       *int     `json:"store_item_id" db:"store_item_id"`
	PurchaseQuantity  *float64 `json:"purchase_quantity" db:"purchase_quantity"`     // null is allowed (e.g., "Bread")
	Note              *string  `json:"note" db:"note"`                               // Optional user note (e.g., "Green bananas")
	CustomItemName    *string  `json:"custom_item_name" db:"custom_item_name"`       // Optional custom item name to allow for shopping list items that don't correspond to a store item
	ShoppingListOrder int      `json:"shopping_list_order" db:"shopping_list_order"` // Order of items in the shopping list
}

type ShoppingList struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
