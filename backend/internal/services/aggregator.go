package services

import (
	"sort"
	"time"

	"github.com/GimhaniHM/backend/internal/models"
	"github.com/GimhaniHM/backend/internal/utils"
)

// holds all transaction data in memory for processing
type Aggregator struct {
	transactions []models.Transaction
}

// creates a Aggregator instance
func NewTestAggregator(txs []models.Transaction) *Aggregator {
	return &Aggregator{transactions: txs}
}

// NewAggregator loads the CSV on startup.
func NewAggregator(csvPath string) (*Aggregator, error) {
	txs, err := utils.ReadTransactions(csvPath)
	if err != nil {
		return nil, err
	}
	return &Aggregator{transactions: txs}, nil
}

// RevenueByCountryAndProduct returns a list of total revenue and transaction count
// grouped by country and product, sorted by highest revenue.
func (a *Aggregator) RevenueByCountryAndProduct() []models.CountryRevenue {
	tmp := map[struct{ C, P string }]struct {
		rev float64
		cnt int
	}{}
	for _, t := range a.transactions {
		key := struct{ C, P string }{t.Country, t.ProductName}
		v := tmp[key]
		v.rev += t.TotalPrice
		v.cnt++
		tmp[key] = v
	}

	out := make([]models.CountryRevenue, 0, len(tmp))
	for k, v := range tmp {
		out = append(out, models.CountryRevenue{
			Country:          k.C,
			ProductName:      k.P,
			TotalRevenue:     v.rev,
			TransactionCount: v.cnt,
		})
	}

	// Sort by descending revenue
	sort.Slice(out, func(i, j int) bool {
		return out[i].TotalRevenue > out[j].TotalRevenue
	})
	return out
}

// TopProducts returns the top N products by total quantity sold
// If two products have the same quantity, they are sorted by name in descending order
func (a *Aggregator) TopProducts(limit int) []models.ProductFrequency {
	tmp := map[string]struct{ cnt, stock int }{}
	for _, t := range a.transactions {
		v := tmp[t.ProductName]
		v.cnt += t.Quantity
		v.stock = t.StockQuantity
		tmp[t.ProductName] = v
	}

	out := make([]models.ProductFrequency, 0, len(tmp))
	for name, v := range tmp {
		out = append(out, models.ProductFrequency{
			ProductName:   name,
			PurchaseCount: v.cnt,
			StockQuantity: v.stock,
		})
	}

	// Sort by purchase count (desc), then by product name (desc)
	sort.Slice(out, func(i, j int) bool {
		if out[i].PurchaseCount == out[j].PurchaseCount {
			return out[i].ProductName > out[j].ProductName
		}
		return out[i].PurchaseCount > out[j].PurchaseCount
	})

	// Limit the result to the top N products
	if len(out) > limit {
		out = out[:limit]
	}
	return out
}

// MonthlySalesVolume returns the quantity sold in each month, sorted chronologically
func (a *Aggregator) MonthlySalesVolume() []models.MonthlySales {
	tmp := map[string]int{}
	for _, t := range a.transactions {
		month := t.TransactionDate.Format("2006-01")
		tmp[month] += t.Quantity
	}

	out := make([]models.MonthlySales, 0, len(tmp))
	for m, vol := range tmp {
		out = append(out, models.MonthlySales{Month: m, SalesVolume: vol})
	}

	// Sort by date (ascending)
	sort.Slice(out, func(i, j int) bool {
		ti, _ := time.Parse("2006-01", out[i].Month)
		tj, _ := time.Parse("2006-01", out[j].Month)
		return ti.Before(tj)
	})
	return out
}

// TopRegionsByRevenue returns the top N regions by total revenue
func (a *Aggregator) TopRegionsByRevenue(limit int) []models.RegionRevenue {
	tmp := map[string]struct {
		rev  float64
		sold int
	}{}
	for _, t := range a.transactions {
		v := tmp[t.Region]
		v.rev += t.TotalPrice
		v.sold += t.Quantity
		tmp[t.Region] = v
	}

	out := make([]models.RegionRevenue, 0, len(tmp))
	for region, v := range tmp {
		out = append(out, models.RegionRevenue{
			Region:       region,
			TotalRevenue: v.rev,
			ItemsSold:    v.sold,
		})
	}

	// Sort by total revenue (desc)
	sort.Slice(out, func(i, j int) bool {
		return out[i].TotalRevenue > out[j].TotalRevenue
	})

	// Limit the result to top N regions
	if len(out) > limit {
		out = out[:limit]
	}
	return out
}
