package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
	"github.com/rjmcnamara10/forage/db/sqlc"
)

// Request/Response DTOs
type CreateShoppingListItemRequest struct {
	ShoppingListID    int32   `json:"shopping_list_id" binding:"required"`
	StoreItemID       *int32  `json:"store_item_id"`
	PurchaseQuantity  *string `json:"purchase_quantity"`
	Note              *string `json:"note"`
	CustomItemName    *string `json:"custom_item_name"`
	ShoppingListOrder int32   `json:"shopping_list_order" binding:"required"`
}

type UpdateShoppingListItemRequest struct {
	StoreItemID       *int32  `json:"store_item_id"`
	PurchaseQuantity  *string `json:"purchase_quantity"`
	Note              *string `json:"note"`
	CustomItemName    *string `json:"custom_item_name"`
	ShoppingListOrder int32   `json:"shopping_list_order" binding:"required"`
}

type ShoppingListItemResponse struct {
	ID                int32   `json:"id"`
	ShoppingListID    int32   `json:"shopping_list_id"`
	StoreItemID       *int32  `json:"store_item_id"`
	PurchaseQuantity  *string `json:"purchase_quantity"`
	Note              *string `json:"note"`
	CustomItemName    *string `json:"custom_item_name"`
	ShoppingListOrder int32   `json:"shopping_list_order"`
}

// GET /shopping-list-items
func getShoppingListItems(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		listID := c.Query("shopping_list_id")
		if listID == "" {
			c.JSON(400, gin.H{"error": "shopping_list_id query parameter is required"})
			return
		}

		shoppingListID, err := strconv.ParseInt(listID, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid shopping list ID"})
			return
		}

		// Verify shopping list exists
		_, err = repos.ShoppingLists.GetShoppingList(c.Request.Context(), int32(shoppingListID))
		if err != nil {
			c.JSON(404, gin.H{"error": "Shopping list not found"})
			return
		}

		items, err := repos.ShoppingListItems.ListShoppingListItems(c.Request.Context(), int32(shoppingListID))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch shopping list items"})
			return
		}

		response := make([]ShoppingListItemResponse, len(items))
		for i, item := range items {
			response[i] = shoppingListItemToResponse(item)
		}

		c.JSON(200, response)
	}
}

// GET /shopping-list-items/:id
func getShoppingListItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid shopping list item ID"})
			return
		}

		item, err := repos.ShoppingListItems.GetShoppingListItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Shopping list item not found"})
			return
		}

		resp := shoppingListItemToResponse(item)
		c.JSON(200, resp)
	}
}

// POST /shopping-list-items
func createShoppingListItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateShoppingListItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Validate that shopping list exists
		_, err := repos.ShoppingLists.GetShoppingList(c.Request.Context(), req.ShoppingListID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Shopping list does not exist"})
			return
		}

		// Validate XOR: exactly one of storeItemID or customItemName must be provided
		hasStoreItemID := req.StoreItemID != nil
		hasCustomItemName := req.CustomItemName != nil && *req.CustomItemName != ""

		if (hasStoreItemID && hasCustomItemName) || (!hasStoreItemID && !hasCustomItemName) {
			c.JSON(400, gin.H{"error": "Exactly one of store_item_id or custom_item_name must be provided"})
			return
		}

		// Convert optional fields to sql.Null types
		storeItemID := sql.NullInt32{Valid: false}
		if req.StoreItemID != nil {
			storeItemID = sql.NullInt32{Int32: *req.StoreItemID, Valid: true}
		}

		purchaseQuantity := sql.NullString{Valid: false}
		if req.PurchaseQuantity != nil {
			purchaseQuantity = sql.NullString{String: *req.PurchaseQuantity, Valid: true}
		}

		note := sql.NullString{Valid: false}
		if req.Note != nil {
			note = sql.NullString{String: *req.Note, Valid: true}
		}

		customItemName := sql.NullString{Valid: false}
		if req.CustomItemName != nil {
			customItemName = sql.NullString{String: *req.CustomItemName, Valid: true}
		}

		params := sqlc.CreateShoppingListItemParams{
			ShoppingListID:    req.ShoppingListID,
			StoreItemID:       storeItemID,
			PurchaseQuantity:  purchaseQuantity,
			Note:              note,
			CustomItemName:    customItemName,
			ShoppingListOrder: req.ShoppingListOrder,
		}

		item, err := repos.ShoppingListItems.CreateShoppingListItem(c.Request.Context(), params)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create shopping list item"})
			return
		}

		resp := shoppingListItemToResponse(item)
		c.JSON(201, resp)
	}
}

// PUT /shopping-list-items/:id
func updateShoppingListItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid shopping list item ID"})
			return
		}

		var req UpdateShoppingListItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Get the existing shopping list item
		existingItem, err := repos.ShoppingListItems.GetShoppingListItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Shopping list item not found"})
			return
		}

		// Validate XOR: exactly one of storeItemID or customItemName must be provided
		hasStoreItemID := req.StoreItemID != nil
		hasCustomItemName := req.CustomItemName != nil && *req.CustomItemName != ""

		if (hasStoreItemID && hasCustomItemName) || (!hasStoreItemID && !hasCustomItemName) {
			c.JSON(400, gin.H{"error": "Exactly one of store_item_id or custom_item_name must be provided"})
			return
		}

		// Convert optional fields, using existing values as defaults
		storeItemID := existingItem.StoreItemID
		if req.StoreItemID != nil {
			storeItemID = sql.NullInt32{Int32: *req.StoreItemID, Valid: true}
		} else if hasCustomItemName {
			// If setting custom item name, clear store item ID
			storeItemID = sql.NullInt32{Valid: false}
		}

		purchaseQuantity := existingItem.PurchaseQuantity
		if req.PurchaseQuantity != nil {
			purchaseQuantity = sql.NullString{String: *req.PurchaseQuantity, Valid: true}
		}

		note := existingItem.Note
		if req.Note != nil {
			note = sql.NullString{String: *req.Note, Valid: true}
		}

		customItemName := existingItem.CustomItemName
		if req.CustomItemName != nil {
			customItemName = sql.NullString{String: *req.CustomItemName, Valid: true}
		} else if hasStoreItemID {
			// If setting store item ID, clear custom item name
			customItemName = sql.NullString{Valid: false}
		}

		params := sqlc.UpdateShoppingListItemParams{
			ID:                int32(id),
			StoreItemID:       storeItemID,
			PurchaseQuantity:  purchaseQuantity,
			Note:              note,
			CustomItemName:    customItemName,
			ShoppingListOrder: req.ShoppingListOrder,
		}

		item, err := repos.ShoppingListItems.UpdateShoppingListItem(c.Request.Context(), params)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update shopping list item"})
			return
		}

		resp := shoppingListItemToResponse(item)
		c.JSON(200, resp)
	}
}

// DELETE /shopping-list-items/:id
func deleteShoppingListItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid shopping list item ID"})
			return
		}

		err = repos.ShoppingListItems.DeleteShoppingListItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete shopping list item"})
			return
		}

		c.JSON(204, nil)
	}
}

// Helper function to convert SQLC ShoppingListItem to response DTO
func shoppingListItemToResponse(item sqlc.ShoppingListItem) ShoppingListItemResponse {
	resp := ShoppingListItemResponse{
		ID:                item.ID,
		ShoppingListID:    item.ShoppingListID,
		StoreItemID:       nil,
		PurchaseQuantity:  nil,
		Note:              nil,
		CustomItemName:    nil,
		ShoppingListOrder: item.ShoppingListOrder,
	}

	if item.StoreItemID.Valid {
		resp.StoreItemID = &item.StoreItemID.Int32
	}
	if item.PurchaseQuantity.Valid {
		resp.PurchaseQuantity = &item.PurchaseQuantity.String
	}
	if item.Note.Valid {
		resp.Note = &item.Note.String
	}
	if item.CustomItemName.Valid {
		resp.CustomItemName = &item.CustomItemName.String
	}

	return resp
}
