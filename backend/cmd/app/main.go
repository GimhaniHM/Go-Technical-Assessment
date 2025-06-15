package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/GimhaniHM/backend/internal/handlers"
	"github.com/GimhaniHM/backend/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// flags for CSV path and listen address
	csvPath := flag.String("data", "data/GO_test_5m.csv", "Path to transactions CSV file")
	addr := flag.String("addr", ":8090", "HTTP listen address")
	workers := flag.Int("workers", runtime.NumCPU(), "Number of CSV parse workers")
	flag.Parse()

	// Run concurrent aggregation
	ca := services.NewConcurrentAggregator(*csvPath, *workers)
	insights, err := ca.Run()
	if err != nil {
		log.Fatalf("aggregation error: %v", err)
	}

	// HTTP handlers
	h := handlers.NewInsightHandler(insights)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/revenue/countries", h.GetCountryRevenue)
		api.GET("/products/top", h.GetTopProducts)
		api.GET("/sales/monthly", h.GetMonthlySales)
		api.GET("/regions/top", h.GetTopRegions)
	}

	log.Printf("Listening on %s", *addr)
	if err := router.Run(*addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
