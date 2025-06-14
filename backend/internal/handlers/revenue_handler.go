package handlers

import (
	"net/http"

	"github.com/GimhaniHM/backend/internal/services"
	"github.com/gin-gonic/gin"
)

// RevenueHandler exposes aggregated endpoints.
type RevenueHandler struct {
	agg *services.AggregateResults
}

// NewRevenueHandler constructs the handler.
func NewRevenueHandler(agg *services.AggregateResults) *RevenueHandler {
	return &RevenueHandler{agg: agg}
}

// GetCountryRevenue handles GET /api/revenue/countries
func (h *RevenueHandler) GetCountryRevenue(c *gin.Context) {
	data, err := h.agg.CountryRevenue()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetTopProducts handles GET /api/products/top
func (h *RevenueHandler) GetTopProducts(c *gin.Context) {
	data, err := h.agg.TopProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetMonthlySales handles GET /api/sales/monthly
func (h *RevenueHandler) GetMonthlySales(c *gin.Context) {
	data, err := h.agg.MonthlySales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// GetTopRegions handles GET /api/regions/top
func (h *RevenueHandler) GetTopRegions(c *gin.Context) {
	data, err := h.agg.TopRegions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
