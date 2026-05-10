package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
)

// Request/Response DTOs
type CreateInventoryItemRequest struct {
	ItemID       int32   `json:"item_id" binding:"required"`
	UnitID       *int32  `json:"unit_id"`
	StoredAmount *string `json:"stored_amount"`
}

type UpdateInventoryItemRequest struct {
	UnitID       *int32  `json:"unit_id"`
	StoredAmount *string `json:"stored_amount"`
}

type InventoryItemResponse struct {
	ID           int32   `json:"id"`
	ItemID       int32   `json:"item_id"`
	UnitID       *int32  `json:"unit_id"`
	StoredAmount *string `json:"stored_amount"`
}

// GET /inventory-items
func getInventoryItems(repos *repository.Repositories) gin.HandlerFunc {
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

		items, err := repos.Inventory.ListInventoryItems(c.Request.Context(), limit, offset)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch inventory items"})
			return
		}

		response := make([]InventoryItemResponse, len(items))
		for i, item := range items {
			response[i] = InventoryItemResponse{
				ID:           item.ID,
				ItemID:       item.ItemID,
				UnitID:       nil,
				StoredAmount: nil,
			}
			if item.UnitID.Valid {
				response[i].UnitID = &item.UnitID.Int32
			}
			if item.StoredAmount.Valid {
				response[i].StoredAmount = &item.StoredAmount.String
			}
		}

		c.JSON(200, response)
	}
}

// GET /inventory-items/:id
func getInventoryItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid inventory item ID"})
			return
		}

		item, err := repos.Inventory.GetInventoryItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Inventory item not found"})
			return
		}

		resp := InventoryItemResponse{
			ID:           item.ID,
			ItemID:       item.ItemID,
			UnitID:       nil,
			StoredAmount: nil,
		}
		if item.UnitID.Valid {
			resp.UnitID = &item.UnitID.Int32
		}
		if item.StoredAmount.Valid {
			resp.StoredAmount = &item.StoredAmount.String
		}

		c.JSON(200, resp)
	}
}

// POST /inventory-items
func createInventoryItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateInventoryItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Validate that item exists
		_, err := repos.Items.GetItem(c.Request.Context(), req.ItemID)
		if err != nil {
			c.JSON(400, gin.H{"error": "Item does not exist"})
			return
		}

		// Check if inventory item already exists for this item
		_, err = repos.Inventory.GetInventoryItemByItemId(c.Request.Context(), req.ItemID)
		if err == nil {
			c.JSON(400, gin.H{"error": "Inventory item already exists for this item"})
			return
		}

		// Convert optional fields to sql.Null types
		unitID := sql.NullInt32{Valid: false}
		if req.UnitID != nil {
			unitID = sql.NullInt32{Int32: *req.UnitID, Valid: true}
		}

		storedAmount := sql.NullString{Valid: false}
		if req.StoredAmount != nil {
			storedAmount = sql.NullString{String: *req.StoredAmount, Valid: true}
		}

		item, err := repos.Inventory.CreateInventoryItem(c.Request.Context(), req.ItemID, unitID, storedAmount)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create inventory item"})
			return
		}

		resp := InventoryItemResponse{
			ID:           item.ID,
			ItemID:       item.ItemID,
			UnitID:       nil,
			StoredAmount: nil,
		}
		if item.UnitID.Valid {
			resp.UnitID = &item.UnitID.Int32
		}
		if item.StoredAmount.Valid {
			resp.StoredAmount = &item.StoredAmount.String
		}

		c.JSON(201, resp)
	}
}

// PUT /inventory-items/:id
func updateInventoryItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid inventory item ID"})
			return
		}

		var req UpdateInventoryItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Get the existing inventory item to retrieve the itemID
		existingItem, err := repos.Inventory.GetInventoryItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Inventory item not found"})
			return
		}

		// Convert optional fields to sql.Null types
		unitID := existingItem.UnitID
		if req.UnitID != nil {
			unitID = sql.NullInt32{Int32: *req.UnitID, Valid: true}
		}

		storedAmount := existingItem.StoredAmount
		if req.StoredAmount != nil {
			storedAmount = sql.NullString{String: *req.StoredAmount, Valid: true}
		}

		item, err := repos.Inventory.UpdateInventoryItem(c.Request.Context(), existingItem.ItemID, unitID, storedAmount)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update inventory item"})
			return
		}

		resp := InventoryItemResponse{
			ID:           item.ID,
			ItemID:       item.ItemID,
			UnitID:       nil,
			StoredAmount: nil,
		}
		if item.UnitID.Valid {
			resp.UnitID = &item.UnitID.Int32
		}
		if item.StoredAmount.Valid {
			resp.StoredAmount = &item.StoredAmount.String
		}

		c.JSON(200, resp)
	}
}

// DELETE /inventory-items/:id
func deleteInventoryItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid inventory item ID"})
			return
		}

		// Get the existing inventory item to retrieve the itemID
		existingItem, err := repos.Inventory.GetInventoryItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Inventory item not found"})
			return
		}

		err = repos.Inventory.DeleteInventoryItem(c.Request.Context(), existingItem.ItemID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete inventory item"})
			return
		}

		c.JSON(204, nil)
	}
}
