package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GimhaniHM/backend/internal/services"
)

func RegisterRevenueRoutes(mux *http.ServeMux, csvPath string) {
	mux.HandleFunc("/api/country-revenue", func(w http.ResponseWriter, r *http.Request) {
		res, err := services.ComputeAggregates(csvPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(res.CountryProducts)
	})

	mux.HandleFunc("/api/top-products", func(w http.ResponseWriter, r *http.Request) {
		res, err := services.ComputeAggregates(csvPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(res.TopProducts)
	})

	mux.HandleFunc("/api/month-sales", func(w http.ResponseWriter, r *http.Request) {
		res, err := services.ComputeAggregates(csvPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(res.MonthlySales)
	})

	mux.HandleFunc("/api/region-revenue", func(w http.ResponseWriter, r *http.Request) {
		res, err := services.ComputeAggregates(csvPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(res.TopRegions)
	})
}
