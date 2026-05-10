package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
)

func RegisterRoutes(router *gin.Engine, repos *repository.Repositories) {
	// Item routes
	router.GET("/items", getItems(repos))
	router.GET("/items/:id", getItem(repos))
	router.POST("/items", createItem(repos))
	router.PUT("/items/:id", updateItem(repos))
	router.DELETE("/items/:id", deleteItem(repos))

	router.GET("/item-categories", getItemCategories(repos))
	router.GET("/item-categories/:id", getItemCategory(repos))
	router.POST("/item-categories", createItemCategory(repos))
	router.PUT("/item-categories/:id", updateItemCategory(repos))
	router.DELETE("/item-categories/:id", deleteItemCategory(repos))

	router.GET("/inventory-items", getInventoryItems(repos))
	router.GET("/inventory-items/:id", getInventoryItem(repos))
	router.POST("/inventory-items", createInventoryItem(repos))
	router.PUT("/inventory-items/:id", updateInventoryItem(repos))
	router.DELETE("/inventory-items/:id", deleteInventoryItem(repos))

	// Unit routes
	router.GET("/units", getUnits(repos))
	router.GET("/units/:id", getUnit(repos))
	router.POST("/units", createUnit(repos))
	router.PUT("/units/:id", updateUnit(repos))
	router.DELETE("/units/:id", deleteUnit(repos))

	// Meal routes
	router.GET("/meals", getMeals(repos))
	router.GET("/meals/:id", getMeal(repos))
	router.POST("/meals", createMeal(repos))
	router.PUT("/meals/:id", updateMeal(repos))
	router.DELETE("/meals/:id", deleteMeal(repos))

	router.GET("/meal-categories", getMealCategories(repos))
	router.GET("/meal-categories/:id", getMealCategory(repos))
	router.POST("/meal-categories", createMealCategory(repos))
	router.PUT("/meal-categories/:id", updateMealCategory(repos))
	router.DELETE("/meal-categories/:id", deleteMealCategory(repos))

	router.GET("/meal-ingredients", getMealIngredients(repos))
	router.GET("/meal-ingredients/:id", getMealIngredient(repos))
	router.POST("/meal-ingredients", createMealIngredient(repos))
	router.PUT("/meal-ingredients/:id", updateMealIngredient(repos))
	router.DELETE("/meal-ingredients/:id", deleteMealIngredient(repos))

	// Store routes
	router.GET("/stores", getStores(repos))
	router.GET("/stores/:id", getStore(repos))
	router.POST("/stores", createStore(repos))
	router.PUT("/stores/:id", updateStore(repos))
	router.DELETE("/stores/:id", deleteStore(repos))

	router.GET("/store-items", getStoreItems(repos))
	router.GET("/store-items/:id", getStoreItem(repos))
	router.POST("/store-items", createStoreItem(repos))
	router.PUT("/store-items/:id", updateStoreItem(repos))
	router.DELETE("/store-items/:id", deleteStoreItem(repos))

	// Shopping List routes
	router.GET("/shopping-lists", getShoppingLists(repos))
	router.GET("/shopping-lists/:id", getShoppingList(repos))
	router.POST("/shopping-lists", createShoppingList(repos))
	router.PUT("/shopping-lists/:id", updateShoppingList(repos))
	router.DELETE("/shopping-lists/:id", deleteShoppingList(repos))

	router.GET("/shopping-list-items", getShoppingListItems(repos))
	router.GET("/shopping-list-items/:id", getShoppingListItem(repos))
	router.POST("/shopping-list-items", createShoppingListItem(repos))
	router.PUT("/shopping-list-items/:id", updateShoppingListItem(repos))
	router.DELETE("/shopping-list-items/:id", deleteShoppingListItem(repos))
}
