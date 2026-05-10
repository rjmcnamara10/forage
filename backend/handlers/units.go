package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rjmcnamara10/forage/db/repository"
)

// Request/Response DTOs
type CreateUnitRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUnitRequest struct {
	Name string `json:"name" binding:"required"`
}

type UnitResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// Pagination response envelope
type PaginatedUnitResponse struct {
	Data   interface{} `json:"data"`
	Total  int64       `json:"total"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

// GET /units
func getUnits(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := int32(50)
		offset := int32(0)
		sortOrder := "ASC"

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

		units, err := repos.Units.ListUnits(c.Request.Context(), limit, offset, sortOrder)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch units"})
			return
		}

		response := make([]UnitResponse, len(units))
		for i, unit := range units {
			response[i] = UnitResponse{
				ID:   unit.ID,
				Name: unit.Name,
			}
		}

		c.JSON(200, response)
	}
}

// GET /units/:id
func getUnit(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid unit ID"})
			return
		}

		unit, err := repos.Units.GetUnit(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "Unit not found"})
			return
		}

		c.JSON(200, UnitResponse{
			ID:   unit.ID,
			Name: unit.Name,
		})
	}
}

// POST /units
func createUnit(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUnitRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		unit, err := repos.Units.CreateUnit(c.Request.Context(), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create unit"})
			return
		}

		c.JSON(201, UnitResponse{
			ID:   unit.ID,
			Name: unit.Name,
		})
	}
}

// PUT /units/:id
func updateUnit(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid unit ID"})
			return
		}

		var req UpdateUnitRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		unit, err := repos.Units.UpdateUnit(c.Request.Context(), int32(id), req.Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update unit"})
			return
		}

		c.JSON(200, UnitResponse{
			ID:   unit.ID,
			Name: unit.Name,
		})
	}
}

// DELETE /units/:id
func deleteUnit(repos *repository.Repositories) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid unit ID"})
			return
		}

		err = repos.Units.DeleteUnit(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete unit"})
			return
		}

		c.JSON(204, nil)
	}
}
