package models

import "time"

type Transaction struct {
	TransactionID   string    // T8d0cd31067f0…
	TransactionDate time.Time // parsed from transaction_date
	UserID          string    // U59971…
	Country         string
	Region          string
	ProductID       string  // P1399820…
	ProductName     string  // Product_399820
	Category        string  // e.g. Toys
	Price           float64 // unit price
	Quantity        int
	TotalPrice      float64 // price * quantity
	StockQuantity   int
	AddedDate       time.Time // parsed from added_date
}

type CountryRevenue struct {
	Country          string  `json:"country"`
	ProductName      string  `json:"product_name"`
	TotalRevenue     float64 `json:"total_revenue"`
	TransactionCount int     `json:"transaction_count"`
}

type ProductFrequency struct {
	ProductName   string `json:"product_name"`
	PurchaseCount int    `json:"purchase_count"`
	StockQuantity int    `json:"stock_quantity"`
}

type MonthlySales struct {
	Month       string  `json:"month"`
	SalesVolume float64 `json:"sales_volume"`
}

type RegionRevenue struct {
	Region       string  `json:"region"`
	TotalRevenue float64 `json:"total_revenue"`
	ItemsSold    int     `json:"items_sold"`
}
