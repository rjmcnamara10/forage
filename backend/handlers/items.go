package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
)

// Request/Response DTOs
type CreateItemRequest struct {
	Name string `json:"name" binding:"required"`
}

type ItemResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// GET /items
func getItems(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := int32(50) // default limit=50
		offset := int32(0) // default offset=0

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

		items, err := repos.Items.ListItems(c.Request.Context(), limit, offset)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch items"})
			return
		}

		response := make([]ItemResponse, len(items))
		for i, item := range items {
			response[i] = ItemResponse{
				ID:   item.ID,
				Name: item.Name,
			}
		}

		c.JSON(200, response)
	}
}

// GET /items/:id
func getItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid item ID"})
			return
		}

		item, err := repos.Items.GetItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Item not found"})
			return
		}

		c.JSON(200, ItemResponse{
			ID:   item.ID,
			Name: item.Name,
		})
	}
}

// POST /items
func createItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		item, err := repos.Items.CreateItem(c.Request.Context(), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create item"})
			return
		}

		c.JSON(201, ItemResponse{
			ID:   item.ID,
			Name: item.Name,
		})
	}
}

// PUT /items/:id
func updateItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid item ID"})
			return
		}

		var req CreateItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		item, err := repos.Items.UpdateItem(c.Request.Context(), int32(id), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update item"})
			return
		}

		c.JSON(200, ItemResponse{
			ID:   item.ID,
			Name: item.Name,
		})
	}
}

// DELETE /items/:id
func deleteItem(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid item ID"})
			return
		}

		err = repos.Items.DeleteItem(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete item"})
			return
		}

		c.JSON(204, nil)
	}
}
