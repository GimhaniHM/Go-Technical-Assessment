package handlers

import (
	"net/http"
	"strconv"

	"github.com/GimhaniHM/backend/internal/services"
	"github.com/gin-gonic/gin"
)

// handles HTTP requests for precomputed insights with optional pagination
type InsightHandler struct {
	data services.Insights
}

// creates a new handler with the given insights
func NewInsightHandler(ins services.Insights) *InsightHandler {
	return &InsightHandler{data: ins}
}

// GetCountryRevenue handles GET requests to return paginated country revenue data.
// Query parameters:
// - limit: number of records to return (default 100)
// - offset: starting position in the dataset (default 0)
func (h *InsightHandler) GetCountryRevenue(c *gin.Context) {
	all := h.data.CountryRevenue

	// parse pagination params
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if err != nil || limit < 1 {
		limit = 100
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}

	// Ensure offset and end index are within bounds
	total := len(all)
	if offset > total {
		offset = total
	}
	end := offset + limit
	if end > total {
		end = total
	}

	// Return paginated data and total count
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  all[offset:end],
	})
}

// GetTopProducts returns the top-20 most purchased products.
func (h *InsightHandler) GetTopProducts(c *gin.Context) {
	c.JSON(http.StatusOK, h.data.TopProducts)
}

// GetMonthlySales returns monthly sales volumes.
func (h *InsightHandler) GetMonthlySales(c *gin.Context) {
	c.JSON(http.StatusOK, h.data.MonthlySales)
}

// GetTopRegions returns the top-30 regions by revenue.
func (h *InsightHandler) GetTopRegions(c *gin.Context) {
	c.JSON(http.StatusOK, h.data.RegionRevenue)
}
