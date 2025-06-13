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
	CountryProducts []models.CountryProd
	TopProducts     []models.ProdStock
	MonthlySales    []models.MonthSales
	TopRegions      []models.RegionAgg
}

func ComputeAggregates(csvPath string) (*AggregateResults, error) {
	jobs := make(chan []string, 1000)
	errs := make(chan error, 1)
	go utils.StreamCSV(csvPath, jobs, errs)

	countryProd := make(map[string]models.CountryProd)
	prodFreq := make(map[string]models.ProdStock)
	monthAgg := make(map[string]models.MonthSales)
	regionAgg := make(map[string]models.RegionAgg)

	var wg sync.WaitGroup
	workerCount := 4
	wg.Add(workerCount)

	for w := 0; w < workerCount; w++ {
		go func() {
			defer wg.Done()
			for rec := range jobs {
				country := rec[3]
				product := rec[6]
				totalPrice, _ := strconv.ParseFloat(rec[10], 64)
				quantity, _ := strconv.Atoi(rec[9])
				stockQty, _ := strconv.Atoi(rec[11])
				region := rec[4]
				month := rec[2][:7]

				// aggregates
				keyCP := country + "|" + product
				cp := countryProd[keyCP]
				cp.Country, cp.ProductName = country, product
				cp.TotalRevenue += totalPrice
				cp.TransactionCount++
				countryProd[keyCP] = cp

				p := prodFreq[product]
				p.ProductName = product
				p.Transactions++
				p.StockQty = stockQty
				prodFreq[product] = p

				m := monthAgg[month]
				m.Month = month
				m.SalesVolume += totalPrice
				monthAgg[month] = m

				r := regionAgg[region]
				r.Region = region
				r.Revenue += totalPrice
				r.ItemsSold += quantity
				regionAgg[region] = r
			}
		}()
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	if err := <-errs; err != nil {
		return nil, err
	}

	// Convert maps to sorted slices
	cpSlice := []models.CountryProd{}
	for _, v := range countryProd {
		cpSlice = append(cpSlice, v)
	}
	sort.Slice(cpSlice, func(i, j int) bool {
		return cpSlice[i].TotalRevenue > cpSlice[j].TotalRevenue
	})

	prodSlice := []models.ProdStock{}
	for _, v := range prodFreq {
		prodSlice = append(prodSlice, v)
	}
	sort.Slice(prodSlice, func(i, j int) bool {
		return prodSlice[i].Transactions > prodSlice[j].Transactions
	})
	if len(prodSlice) > 20 {
		prodSlice = prodSlice[:20]
	}

	monthSlice := []models.MonthSales{}
	for _, v := range monthAgg {
		monthSlice = append(monthSlice, v)
	}
	sort.Slice(monthSlice, func(i, j int) bool {
		return monthSlice[i].SalesVolume > monthSlice[j].SalesVolume
	})

	regionSlice := []models.RegionAgg{}
	for _, v := range regionAgg {
		regionSlice = append(regionSlice, v)
	}
	sort.Slice(regionSlice, func(i, j int) bool {
		return regionSlice[i].Revenue > regionSlice[j].Revenue
	})
	if len(regionSlice) > 30 {
		regionSlice = regionSlice[:30]
	}

	return &AggregateResults{cpSlice, prodSlice, monthSlice, regionSlice}, nil
}
