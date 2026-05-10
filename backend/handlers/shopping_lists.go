package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
)

// Request/Response DTOs
type CreateShoppingListRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateShoppingListRequest struct {
	Name string `json:"name" binding:"required"`
}

type ShoppingListResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// GET /shopping-lists
func getShoppingLists(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := int32(50)
		offset := int32(0)
		sortOrder := "DESC" // Default: most recent first

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

		if so := c.Query("sort_order"); so != "" && (so == "ASC" || so == "DESC") {
			sortOrder = so
		}

		shoppingLists, err := repos.ShoppingLists.ListShoppingLists(c.Request.Context(), limit, offset, sortOrder)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch shopping lists"})
			return
		}

		response := make([]ShoppingListResponse, len(shoppingLists))
		for i, sl := range shoppingLists {
			response[i] = ShoppingListResponse{
				ID:        sl.ID,
				Name:      sl.Name,
				CreatedAt: sl.CreatedAt,
			}
		}

		total, _ := repos.ShoppingLists.ListShoppingListsCount(c.Request.Context())

		c.JSON(200, PaginatedResponse{
			Data:   response,
			Total:  total,
			Limit:  limit,
			Offset: offset,
		})
	}
}

// GET /shopping-lists/:id
func getShoppingList(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid shopping list ID"})
			return
		}

		shoppingList, err := repos.ShoppingLists.GetShoppingList(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Shopping list not found"})
			return
		}

		c.JSON(200, ShoppingListResponse{
			ID:        shoppingList.ID,
			Name:      shoppingList.Name,
			CreatedAt: shoppingList.CreatedAt,
		})
	}
}

// POST /shopping-lists
func createShoppingList(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateShoppingListRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		shoppingList, err := repos.ShoppingLists.CreateShoppingList(c.Request.Context(), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create shopping list"})
			return
		}

		c.JSON(201, ShoppingListResponse{
			ID:        shoppingList.ID,
			Name:      shoppingList.Name,
			CreatedAt: shoppingList.CreatedAt,
		})
	}
}

// PUT /shopping-lists/:id
func updateShoppingList(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid shopping list ID"})
			return
		}

		var req UpdateShoppingListRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		shoppingList, err := repos.ShoppingLists.UpdateShoppingList(c.Request.Context(), int32(id), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update shopping list"})
			return
		}

		c.JSON(200, ShoppingListResponse{
			ID:        shoppingList.ID,
			Name:      shoppingList.Name,
			CreatedAt: shoppingList.CreatedAt,
		})
	}
}

// DELETE /shopping-lists/:id
func deleteShoppingList(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid shopping list ID"})
			return
		}

		err = repos.ShoppingLists.DeleteShoppingList(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete shopping list"})
			return
		}

		c.JSON(204, nil)
	}
}
