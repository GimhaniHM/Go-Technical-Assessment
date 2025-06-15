package services

import (
	"reflect"
	"testing"
	"time"

	"github.com/GimhaniHM/backend/internal/models"
)

// create a dummy transaction
func makeTransaction(country, product string, qty int, price float64) models.Transaction {
	return models.Transaction{
		Country:     country,
		ProductName: product,
		Quantity:    qty,
		Price:       price,
		TotalPrice:  float64(qty) * price,
	}
}

// TestRevenueByCountryAndProduct checks whether revenue is correctly grouped
// by country and product and sorted by total revenue
func TestRevenueByCountryAndProduct(t *testing.T) {
	agg := &Aggregator{
		transactions: []models.Transaction{
			makeTransaction("A", "X", 2, 10.0), // rev=20
			makeTransaction("A", "X", 3, 10.0), // rev=30
			makeTransaction("B", "Y", 1, 5.0),  // rev=5
		},
	}

	got := agg.RevenueByCountryAndProduct()

	want := []models.CountryRevenue{
		{Country: "A", ProductName: "X", TotalRevenue: 50.0, TransactionCount: 2},
		{Country: "B", ProductName: "Y", TotalRevenue: 5.0, TransactionCount: 1},
	}

	// validate the result
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RevenueByCountryAndProduct() = %#v; want %#v", got, want)
	}
}

// TestTopProducts checks whether the most purchased products are correctly
// identified and sorted by quantity (with tie-breaker on name descending)
func TestTopProducts(t *testing.T) {
	agg := &Aggregator{
		transactions: []models.Transaction{
			makeTransaction("", "P1", 3, 1.0),
			makeTransaction("", "P2", 5, 1.0),
			makeTransaction("", "P1", 2, 1.0),
		},
	}

	got := agg.TopProducts(10)
	want := []models.ProductFrequency{
		{ProductName: "P2", PurchaseCount: 5, StockQuantity: 0},
		{ProductName: "P1", PurchaseCount: 5, StockQuantity: 0},
	}

	// validate the result
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopProducts() = %+v; want %+v", got, want)
	}
}

// TestMonthlySalesVolume verifies that monthly sales volume is calculated
// and grouped correctly by year and month
func TestMonthlySalesVolume(t *testing.T) {
	parse := func(s string) time.Time {
		d, _ := time.Parse("2006-01-02", s)
		return d
	}
	agg := &Aggregator{
		transactions: []models.Transaction{
			{TransactionDate: parse("2024-01-01"), Quantity: 2},
			{TransactionDate: parse("2024-01-15"), Quantity: 3},
			{TransactionDate: parse("2024-02-01"), Quantity: 1},
		},
	}
	got := agg.MonthlySalesVolume()

	// Validate the result
	if len(got) != 2 ||
		got[0].Month != "2024-01" || got[0].SalesVolume != 5 ||
		got[1].Month != "2024-02" || got[1].SalesVolume != 1 {
		t.Errorf("MonthlySalesVolume() = %+v; want [{2024-01 5} {2024-02 1}]", got)
	}
}

// TestTopRegionsByRevenue checks that the regions with the highest revenue
// are correctly calculated and sorted in descending order
func TestTopRegionsByRevenue(t *testing.T) {
	agg := &Aggregator{
		transactions: []models.Transaction{
			{Region: "R1", Quantity: 2, TotalPrice: 20.0},
			{Region: "R2", Quantity: 1, TotalPrice: 5.0},
			{Region: "R1", Quantity: 1, TotalPrice: 10.0},
		},
	}

	got := agg.TopRegionsByRevenue(10)
	want := []models.RegionRevenue{
		{Region: "R1", TotalRevenue: 30.0, ItemsSold: 3},
		{Region: "R2", TotalRevenue: 5.0, ItemsSold: 1},
	}

	// Validate the result
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopRegionsByRevenue() = %+v; want %+v", got, want)
	}
}
