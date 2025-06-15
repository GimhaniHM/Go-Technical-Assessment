package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GimhaniHM/backend/internal/models"
	"github.com/GimhaniHM/backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCountryRevenue(t *testing.T) {
	// Prepare dummy aggregator
	mockAgg := services.NewTestAggregator([]models.Transaction{
		{Country: "USA", ProductName: "Prod1", TotalPrice: 100, Quantity: 1},
	})

	// Prepare a new RevenueHandler with the dummy aggregator
	handler := NewRevenueHandler(mockAgg)

	// Set up a Gin router and define the test route
	router := gin.Default()
	router.GET("/api/revenue/countries", handler.GetCountryRevenue)

	// Create a test HTTP request and response recorder
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/revenue/countries", nil)

	router.ServeHTTP(w, req)

	// Validate the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "USA")
}
