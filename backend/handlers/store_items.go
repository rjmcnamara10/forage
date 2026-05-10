package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
	"github.com/rjmcnamara10/forage/db/sqlc"
)

// Request/Response DTOs
type CreateStoreItemRequest struct {
	StoreID                   int32  `json:"store_id" binding:"required"`
	ItemID                    int32  `json:"item_id" binding:"required"`
	PurchaseUnitID            *int32 `json:"purchase_unit_id"`
	InventoryUnitID           *int32 `json:"inventory_unit_id"`
	InventoryUnitsPerPurchase string `json:"inventory_units_per_purchase"`
	PurchasedByDecimal        bool   `json:"purchased_by_decimal"`
	IsFavorite                bool   `json:"is_favorite"`
	StoreTraversalOrder       int32  `json:"store_traversal_order" binding:"required"`
}

type UpdateStoreItemRequest struct {
	PurchaseUnitID            *int32 `json:"purchase_unit_id"`
	InventoryUnitID           *int32 `json:"inventory_unit_id"`
	InventoryUnitsPerPurchase string `json:"inventory_units_per_purchase"`
	PurchasedByDecimal        bool   `json:"purchased_by_decimal"`
	IsFavorite                bool   `json:"is_favorite"`
	StoreTraversalOrder       int32  `json:"store_traversal_order" binding:"required"`
}

type StoreItemResponse struct {
	ID                        int32  `json:"id"`
	StoreID                   int32  `json:"store_id"`
	ItemID                    int32  `json:"item_id"`
	PurchaseUnitID            *int32 `json:"purchase_unit_id"`
	InventoryUnitID           *int32 `json:"inventory_unit_id"`
	InventoryUnitsPerPurchase string `json:"inventory_units_per_purchase"`
	PurchasedByDecimal        bool   `json:"purchased_by_decimal"`
	IsFavorite                bool   `json:"is_favorite"`
	StoreTraversalOrder       int32  `json:"store_traversal_order"`
}

// GET /store-items
func getStoreItems(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := int32(50)
		offset := int32(0)

		if l := c.Query("limit"); l != "" {
			if parsed, err := strconv.ParseInt(l, 10, 32); err == nil {
				limit = int32(parsed)
			}
		}

		if o := c.Query("offset"); o != "" {
			if parsed, err := strconv.ParseInt(o, 10, 32); err == nil {
				offset = int32(parsed)
			}
		}

		items, err := repos.StoreItems.ListStoreItems(c.Request.Context(), limit, offset)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch store items"})
			return
		}

		response := make([]StoreItemResponse, len(items))
		for i, item := range items {
			response[i] = storeItemToResponse(item)
		}

		c.JSON(200, response)
	}
}

// GET /store-items/:id
func getStoreItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid store item ID"})
			return
		}

		item, err := repos.StoreItems.GetStoreItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Store item not found"})
			return
		}

		resp := storeItemToResponse(item)
		c.JSON(200, resp)
	}
}

// GET /store-items?store_id=123
func getStoreItemsByStore(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		storeIDStr := c.Query("store_id")
		if storeIDStr == "" {
			c.JSON(400, gin.H{"error": "store_id query parameter is required"})
			return
		}

		storeID, err := strconv.ParseInt(storeIDStr, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid store ID"})
			return
		}

		// Verify store exists
		_, err = repos.Stores.GetStore(c.Request.Context(), int32(storeID))
		if err != nil {
			c.JSON(404, gin.H{"error": "Store not found"})
			return
		}

		items, err := repos.StoreItems.ListStoreItemsByStore(c.Request.Context(), int32(storeID))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch store items"})
			return
		}

		response := make([]StoreItemResponse, len(items))
		for i, item := range items {
			response[i] = storeItemToResponse(item)
		}

		c.JSON(200, response)
	}
}

// POST /store-items
func createStoreItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateStoreItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Validate that store exists
		_, err := repos.Stores.GetStore(c.Request.Context(), req.StoreID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Store does not exist"})
			return
		}

		// Validate that item exists
		_, err = repos.Items.GetItem(c.Request.Context(), req.ItemID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Item does not exist"})
			return
		}

		// Validate purchasedByDecimal constraint
		if req.PurchasedByDecimal && req.InventoryUnitsPerPurchase != "1.0" && req.InventoryUnitsPerPurchase != "1" {
			c.JSON(400, gin.H{"error": "If purchased_by_decimal is true, inventory_units_per_purchase must be 1.0"})
			return
		}

		// Convert optional fields to sql.Null types
		purchaseUnitID := sql.NullInt32{Valid: false}
		if req.PurchaseUnitID != nil {
			purchaseUnitID = sql.NullInt32{Int32: *req.PurchaseUnitID, Valid: true}
		}

		inventoryUnitID := sql.NullInt32{Valid: false}
		if req.InventoryUnitID != nil {
			inventoryUnitID = sql.NullInt32{Int32: *req.InventoryUnitID, Valid: true}
		}

		params := sqlc.CreateStoreItemParams{
			StoreID:                   req.StoreID,
			ItemID:                    req.ItemID,
			PurchaseUnitID:            purchaseUnitID,
			InventoryUnitID:           inventoryUnitID,
			InventoryUnitsPerPurchase: req.InventoryUnitsPerPurchase,
			PurchasedByDecimal:        req.PurchasedByDecimal,
			IsFavorite:                req.IsFavorite,
			StoreTraversalOrder:       req.StoreTraversalOrder,
		}

		item, err := repos.StoreItems.CreateStoreItem(c.Request.Context(), params)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create store item"})
			return
		}

		resp := storeItemToResponse(item)
		c.JSON(201, resp)
	}
}

// PUT /store-items/:id
func updateStoreItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid store item ID"})
			return
		}

		var req UpdateStoreItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Get the existing store item
		existingItem, err := repos.StoreItems.GetStoreItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Store item not found"})
			return
		}

		// Validate purchasedByDecimal constraint
		if req.PurchasedByDecimal && req.InventoryUnitsPerPurchase != "1.0" && req.InventoryUnitsPerPurchase != "1" {
			c.JSON(400, gin.H{"error": "If purchased_by_decimal is true, inventory_units_per_purchase must be 1.0"})
			return
		}

		// Convert optional fields to sql.Null types
		purchaseUnitID := existingItem.PurchaseUnitID
		if req.PurchaseUnitID != nil {
			purchaseUnitID = sql.NullInt32{Int32: *req.PurchaseUnitID, Valid: true}
		}

		inventoryUnitID := existingItem.InventoryUnitID
		if req.InventoryUnitID != nil {
			inventoryUnitID = sql.NullInt32{Int32: *req.InventoryUnitID, Valid: true}
		}

		params := sqlc.UpdateStoreItemParams{
			ID:                        int32(id),
			PurchaseUnitID:            purchaseUnitID,
			InventoryUnitID:           inventoryUnitID,
			InventoryUnitsPerPurchase: req.InventoryUnitsPerPurchase,
			PurchasedByDecimal:        req.PurchasedByDecimal,
			IsFavorite:                req.IsFavorite,
			StoreTraversalOrder:       req.StoreTraversalOrder,
		}

		item, err := repos.StoreItems.UpdateStoreItem(c.Request.Context(), params)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update store item"})
			return
		}

		resp := storeItemToResponse(item)
		c.JSON(200, resp)
	}
}

// DELETE /store-items/:id
func deleteStoreItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid store item ID"})
			return
		}

		err = repos.StoreItems.DeleteStoreItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete store item"})
			return
		}

		c.JSON(204, nil)
	}
}

// Helper function to convert SQLC StoreItem to response DTO
func storeItemToResponse(item sqlc.StoreItem) StoreItemResponse {
	resp := StoreItemResponse{
		ID:                        item.ID,
		StoreID:                   item.StoreID,
		ItemID:                    item.ItemID,
		PurchaseUnitID:            nil,
		InventoryUnitID:           nil,
		InventoryUnitsPerPurchase: item.InventoryUnitsPerPurchase,
		PurchasedByDecimal:        item.PurchasedByDecimal,
		IsFavorite:                item.IsFavorite,
		StoreTraversalOrder:       item.StoreTraversalOrder,
	}

	if item.PurchaseUnitID.Valid {
		resp.PurchaseUnitID = &item.PurchaseUnitID.Int32
	}
	if item.InventoryUnitID.Valid {
		resp.InventoryUnitID = &item.InventoryUnitID.Int32
	}

	return resp
}
