package models

import "time"

// Transaction represents one row in your CSV file.
type Transaction struct {
	TransactionID   string    // e.g. T8d0cd31067f0
	TransactionDate time.Time // parsed from transaction_date
	UserID          string    // e.g. U59971
	Country         string
	Region          string
	ProductID       string    // e.g. P1399820
	ProductName     string    // e.g. Product_399820
	Category        string    // e.g. Toys
	Price           float64   // unit price
	Quantity        int       // quantity sold
	TotalPrice      float64   // price * quantity
	StockQuantity   int       // current stock quantity
	AddedDate       time.Time // parsed from added_date
}

// CountryRevenue for /api/revenue/countries
type CountryRevenue struct {
	Country          string  `json:"country"`
	ProductName      string  `json:"product_name"`
	TotalRevenue     float64 `json:"total_revenue"`
	TransactionCount int     `json:"transaction_count"`
}

// ProductFrequency for /api/products/top
type ProductFrequency struct {
	ProductName   string `json:"product_name"`
	PurchaseCount int    `json:"purchase_count"`
	StockQuantity int    `json:"stock_quantity"`
}

// MonthlySales for /api/sales/monthly
type MonthlySales struct {
	Month       string `json:"month"`
	SalesVolume int    `json:"sales_volume"`
}

// RegionRevenue for /api/regions/top
type RegionRevenue struct {
	Region       string  `json:"region"`
	TotalRevenue float64 `json:"total_revenue"`
	ItemsSold    int     `json:"items_sold"`
}
