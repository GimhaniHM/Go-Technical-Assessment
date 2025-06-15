package handlers

import (
	"net/http"
	"strconv"

	"github.com/GimhaniHM/backend/internal/services"
	"github.com/gin-gonic/gin"
)

// handles HTTP requests for revenue-related insights
type RevenueHandler struct {
	agg *services.Aggregator
}

// initializes a new RevenueHandler with the given aggregator
func NewRevenueHandler(agg *services.Aggregator) *RevenueHandler {
	return &RevenueHandler{agg: agg}
}

// GetCountryRevenue handles GET requests for country-level revenue data
// Supports pagination using 'limit' and 'offset' query parameters
func (h *RevenueHandler) GetCountryRevenue(c *gin.Context) {
	// Fetch all country revenue data
	all := h.agg.RevenueByCountryAndProduct()

	// parse pagination params
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if err != nil || limit < 1 {
		limit = 100
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}

	// Calculate bounds to avoid overflow
	total := len(all)
	if offset > total {
		offset = total
	}
	end := offset + limit
	if end > total {
		end = total
	}

	page := all[offset:end]

	// wrap with total + data
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  page,
	})
}

// GetTopProducts handles GET requests for the top-N most purchased products
// Accepts 'limit' query parameter (default is 20)
func (h *RevenueHandler) GetTopProducts(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit < 1 {
		limit = 20
	}
	c.JSON(http.StatusOK, h.agg.TopProducts(limit))
}

// GetMonthlySales handles GET requests for monthly sales volume data
func (h *RevenueHandler) GetMonthlySales(c *gin.Context) {
	c.JSON(http.StatusOK, h.agg.MonthlySalesVolume())
}

// GetTopRegions handles GET requests for the top-N regions by total revenue
// Accepts 'limit' query parameter (default is 30)
func (h *RevenueHandler) GetTopRegions(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "30"))
	if err != nil || limit < 1 {
		limit = 30
	}
	c.JSON(http.StatusOK, h.agg.TopRegionsByRevenue(limit))
}
