package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
)

// Request/Response DTOs
type CreateStoreRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateStoreRequest struct {
	Name string `json:"name" binding:"required"`
}

type StoreResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// GET /stores
func getStores(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		stores, err := repos.Stores.ListStores(c.Request.Context())
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch stores"})
			return
		}

		response := make([]StoreResponse, len(stores))
		for i, store := range stores {
			response[i] = StoreResponse{
				ID:   store.ID,
				Name: store.Name,
			}
		}

		c.JSON(200, response)
	}
}

// GET /stores/:id
func getStore(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid store ID"})
			return
		}

		store, err := repos.Stores.GetStore(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Store not found"})
			return
		}

		c.JSON(200, StoreResponse{
			ID:   store.ID,
			Name: store.Name,
		})
	}
}

// POST /stores
func createStore(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateStoreRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		store, err := repos.Stores.CreateStore(c.Request.Context(), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create store"})
			return
		}

		c.JSON(201, StoreResponse{
			ID:   store.ID,
			Name: store.Name,
		})
	}
}

// PUT /stores/:id
func updateStore(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid store ID"})
			return
		}

		var req UpdateStoreRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		store, err := repos.Stores.UpdateStore(c.Request.Context(), int32(id), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update store"})
			return
		}

		c.JSON(200, StoreResponse{
			ID:   store.ID,
			Name: store.Name,
		})
	}
}

// DELETE /stores/:id
func deleteStore(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid store ID"})
			return
		}

		err = repos.Stores.DeleteStore(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete store"})
			return
		}

		c.JSON(204, nil)
	}
}
