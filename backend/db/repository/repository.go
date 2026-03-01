package repository

import (
	"database/sql"
	"github.com/rjmcnamara10/forage/db/sqlc"
)

// Repositories holds all repository instances for dependency injection
type Repositories struct {
	Units             *UnitRepository
	ItemCategories    *ItemCategoryRepository
	MealCategories    *MealCategoryRepository
	Items             *ItemRepository
	Inventory         *InventoryRepository
	Meals             *MealRepository
	MealIngredients   *MealIngredientRepository
	Stores            *StoreRepository
	StoreItems        *StoreItemRepository
	ShoppingLists     *ShoppingListRepository
	ShoppingListItems *ShoppingListItemRepository
	Grocery           *GroceryRepository
}

// NewRepositories creates and returns all repositories
func NewRepositories(db *sql.DB) *Repositories {
	queries := sqlc.New(db)

	return &Repositories{
		Units:             NewUnitRepository(queries),
		ItemCategories:    NewItemCategoryRepository(queries),
		MealCategories:    NewMealCategoryRepository(queries),
		Items:             NewItemRepository(queries),
		Inventory:         NewInventoryRepository(queries),
		Meals:             NewMealRepository(queries),
		MealIngredients:   NewMealIngredientRepository(queries),
		Stores:            NewStoreRepository(queries),
		StoreItems:        NewStoreItemRepository(queries),
		ShoppingLists:     NewShoppingListRepository(queries),
		ShoppingListItems: NewShoppingListItemRepository(queries),
		Grocery:           NewGroceryRepository(queries),
	}
}
