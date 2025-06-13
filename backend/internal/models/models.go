package models

type CountryProd struct {
	Country          string  `json:"country"`
	ProductName      string  `json:"product_name"`
	TotalRevenue     float64 `json:"total_revenue"`
	TransactionCount int     `json:"transaction_count"`
}

type ProdStock struct {
	ProductName  string `json:"product_name"`
	Transactions int    `json:"transactions"`
	StockQty     int    `json:"stock_qty"`
}

type MonthSales struct {
	Month       string  `json:"month"`
	SalesVolume float64 `json:"sales_volume"`
}

type RegionAgg struct {
	Region    string  `json:"region"`
	Revenue   float64 `json:"revenue"`
	ItemsSold int     `json:"items_sold"`
}
