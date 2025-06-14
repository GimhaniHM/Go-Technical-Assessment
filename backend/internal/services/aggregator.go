package services

import (
	"sort"
	"strconv"
	"sync"

	"github.com/GimhaniHM/backend/internal/models"
	"github.com/GimhaniHM/backend/internal/utils"
)

// AggregateResults holds all processed dataset slices
type AggregateResults struct {
	CountryRevenueTable []models.CountryRevenue   // country + product revenue
	TopProducts         []models.ProductFrequency // top 20 products
	MonthlySales        []models.MonthlySales     // monthly sales volumes
	TopRegions          []models.RegionRevenue    // top 30 regions
}

// ComputeAggregates reads the CSV at csvPath, aggregates metrics, and returns sorted slices.
func ComputeAggregates(csvPath string) (*AggregateResults, error) {
	jobs := make(chan []string, 1000)
	errs := make(chan error, 1)
	go utils.StreamCSV(csvPath, jobs, errs)

	// In-memory aggregation maps
	cpMap := make(map[string]models.CountryRevenue)
	pfMap := make(map[string]models.ProductFrequency)
	msMap := make(map[string]models.MonthlySales)
	rrMap := make(map[string]models.RegionRevenue)

	var wg sync.WaitGroup
	workerCount := 4
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for rec := range jobs {
				// parse fields from CSV record
				country := rec[3]
				region := rec[4]
				date := rec[2]    // "YYYY-MM-DD..."
				month := date[:7] // "YYYY-MM"
				product := rec[6]
				totalPrice, _ := strconv.ParseFloat(rec[10], 64)
				quantity, _ := strconv.Atoi(rec[9])
				stockQty, _ := strconv.Atoi(rec[11])

				// 1) Country + Product → CountryRevenue
				keyCP := country + "|" + product
				cr := cpMap[keyCP]
				cr.Country = country
				cr.ProductName = product
				cr.TotalRevenue += totalPrice
				cr.TransactionCount++
				cpMap[keyCP] = cr

				// 2) Product → ProductFrequency
				pf := pfMap[product]
				pf.ProductName = product
				pf.PurchaseCount++
				pf.StockQuantity = stockQty
				pfMap[product] = pf

				// 3) YYYY-MM → MonthlySales
				ms := msMap[month]
				ms.Month = month
				ms.SalesVolume += totalPrice
				msMap[month] = ms

				// 4) Region → RegionRevenue
				rr := rrMap[region]
				rr.Region = region
				rr.TotalRevenue += totalPrice
				rr.ItemsSold += quantity
				rrMap[region] = rr
			}
		}()
	}

	// wait for workers to finish then check errors
	go func() {
		wg.Wait()
		close(errs)
	}()

	if err := <-errs; err != nil {
		return nil, err
	}

	// Convert maps to slices
	var countrySlice []models.CountryRevenue
	for _, v := range cpMap {
		countrySlice = append(countrySlice, v)
	}
	sort.Slice(countrySlice, func(i, j int) bool {
		return countrySlice[i].TotalRevenue > countrySlice[j].TotalRevenue
	})

	var productSlice []models.ProductFrequency
	for _, v := range pfMap {
		productSlice = append(productSlice, v)
	}
	sort.Slice(productSlice, func(i, j int) bool {
		return productSlice[i].PurchaseCount > productSlice[j].PurchaseCount
	})
	if len(productSlice) > 20 {
		productSlice = productSlice[:20]
	}

	var monthSlice []models.MonthlySales
	for _, v := range msMap {
		monthSlice = append(monthSlice, v)
	}
	sort.Slice(monthSlice, func(i, j int) bool {
		return monthSlice[i].SalesVolume > monthSlice[j].SalesVolume
	})

	var regionSlice []models.RegionRevenue
	for _, v := range rrMap {
		regionSlice = append(regionSlice, v)
	}
	sort.Slice(regionSlice, func(i, j int) bool {
		return regionSlice[i].TotalRevenue > regionSlice[j].TotalRevenue
	})
	if len(regionSlice) > 30 {
		regionSlice = regionSlice[:30]
	}

	return &AggregateResults{
		CountryRevenueTable: countrySlice,
		TopProducts:         productSlice,
		MonthlySales:        monthSlice,
		TopRegions:          regionSlice,
	}, nil
}
