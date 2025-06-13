package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/GimhaniHM/backend/internal/handlers"
)

func main() {
	csvPath := filepath.Join("..", "data", "GO_test_5m.csv")
	mux := http.NewServeMux()
	handlers.RegisterRevenueRoutes(mux, csvPath)

	log.Println("Listening on :5000")
	log.Fatal(http.ListenAndServe(":5000", mux))
}
