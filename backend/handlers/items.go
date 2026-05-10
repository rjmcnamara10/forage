package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
)

// Request/Response DTOs
type CreateItemRequest struct {
	Name        string  `json:"name" binding:"required"`
	CategoryIDs []int32 `json:"category_ids" binding:"required,min=1"`
}

type UpdateItemRequest struct {
	Name        string  `json:"name" binding:"required"`
	CategoryIDs []int32 `json:"category_ids" binding:"required,min=1"`
}

type ItemResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// Pagination response envelope
type PaginatedResponse struct {
	Data   interface{} `json:"data"`
	Total  int64       `json:"total"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

// GET /items
func getItems(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := int32(50)
		offset := int32(0)
		sortOrder := "ASC"
		query := c.Query("q")
		categoryID := c.Query("category_id")

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

		var items []ItemResponse
		var total int64

		// Determine which query to use based on search and filter params
		if query != "" {
			// Search mode
			searchItems, err := repos.Items.SearchItems(c.Request.Context(), query, limit, offset, sortOrder)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to fetch items"})
				return
			}
			items = make([]ItemResponse, len(searchItems))
			for i, item := range searchItems {
				items[i] = ItemResponse{
					ID:   item.ID,
					Name: item.Name,
				}
			}
			total, _ = repos.Items.SearchItemsCount(c.Request.Context(), query)
		} else if categoryID != "" {
			// Category filter mode
			catID, err := strconv.ParseInt(categoryID, 10, 32)
			if err != nil {
				c.JSON(400, gin.H{"error": "Invalid category ID"})
				return
			}
			filterItems, err := repos.Items.ListItemsByCategory(c.Request.Context(), int32(catID), limit, offset, sortOrder)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to fetch items"})
				return
			}
			items = make([]ItemResponse, len(filterItems))
			for i, item := range filterItems {
				items[i] = ItemResponse{
					ID:   item.ID,
					Name: item.Name,
				}
			}
			total, _ = repos.Items.ListItemsByCategoryCount(c.Request.Context(), int32(catID))
		} else {
			// List all items
			listItems, err := repos.Items.ListItems(c.Request.Context(), limit, offset, sortOrder)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to fetch items"})
				return
			}
			items = make([]ItemResponse, len(listItems))
			for i, item := range listItems {
				items[i] = ItemResponse{
					ID:   item.ID,
					Name: item.Name,
				}
			}
			total, _ = repos.Items.ListItemsCount(c.Request.Context())
		}

		c.JSON(200, PaginatedResponse{
			Data:   items,
			Total:  total,
			Limit:  limit,
			Offset: offset,
		})
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

		// Validate that at least one category is provided
		if len(req.CategoryIDs) == 0 {
			c.JSON(400, gin.H{"error": "At least one category_id must be provided"})
			return
		}

		// Validate all categories exist
		for _, categoryID := range req.CategoryIDs {
			_, err := repos.ItemCategories.GetItemCategory(c.Request.Context(), categoryID)
			if err != nil {
				c.JSON(400, gin.H{"error": "One or more categories do not exist"})
				return
			}
		}

		item, err := repos.Items.CreateItem(c.Request.Context(), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create item"})
			return
		}

		// Add categories to item
		for _, categoryID := range req.CategoryIDs {
			err := repos.Items.AddItemCategory(c.Request.Context(), item.ID, categoryID)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to add categories to item"})
				return
			}
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

		var req UpdateItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// Validate that at least one category is provided
		if len(req.CategoryIDs) == 0 {
			c.JSON(400, gin.H{"error": "At least one category_id must be provided"})
			return
		}

		// Validate all categories exist
		for _, categoryID := range req.CategoryIDs {
			_, err := repos.ItemCategories.GetItemCategory(c.Request.Context(), categoryID)
			if err != nil {
				c.JSON(400, gin.H{"error": "One or more categories do not exist"})
				return
			}
		}

		item, err := repos.Items.UpdateItem(c.Request.Context(), int32(id), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update item"})
			return
		}

		// Remove all existing categories
		err = repos.Items.RemoveAllItemCategories(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update item categories"})
			return
		}

		// Add new categories
		for _, categoryID := range req.CategoryIDs {
			err := repos.Items.AddItemCategory(c.Request.Context(), int32(id), categoryID)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to add categories to item"})
				return
			}
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
