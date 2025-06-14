package main

import (
	"flag"
	"log"

	"github.com/GimhaniHM/backend/internal/handlers"
	"github.com/GimhaniHM/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// flags for CSV path and listen address
	csvPath := flag.String("data", "data/transactions.csv", "Path to transactions CSV file")
	addr := flag.String("addr", ":8090", "Server listen address")
	flag.Parse()

	// load & aggregrate data
	agg, err := services.ComputeAggregates(*csvPath)
	if err != nil {
		log.Fatalf("failed to load data: %v", err)
	}

	// set up HTTP handlers
	revH := handlers.NewRevenueHandler(agg)
	router := gin.Default()

	// Group all analytics routes under /api
	api := router.Group("/api")
	{
		api.GET("/country-revenue", revH.GetCountryRevenue(csvPath))
		api.GET("/top-products", revH.GetTopProducts(csvPath))
		api.GET("/monthly-sales", revH.GetMonthlySales(csvPath))
		api.GET("/region-revenue", revH.GetTopRegions(csvPath))
	}

	log.Printf("▶️  listening on %s", *addr)
	if err := router.Run(*addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
