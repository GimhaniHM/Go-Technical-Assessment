package services

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/GimhaniHM/backend/internal/models"
)

// handles concurrent processing of large CSV data files
type ConcurrentAggregator struct {
	filePath string
	workers  int
}

// holds the final aggregated results to be returned
type Insights struct {
	CountryRevenue []models.CountryRevenue
	TopProducts    []models.ProductFrequency
	MonthlySales   []models.MonthlySales
	RegionRevenue  []models.RegionRevenue
}

// creates and returns a new ConcurrentAggregator instance
func NewConcurrentAggregator(path string, workers int) *ConcurrentAggregator {
	return &ConcurrentAggregator{filePath: path, workers: workers}
}

// Run reads the CSV, processes it concurrently, aggregates results, and returns insight
func (ca *ConcurrentAggregator) Run() (Insights, error) {
	// Open the CSV file
	f, err := os.Open(ca.filePath)
	if err != nil {
		return Insights{}, err
	}
	defer f.Close()
	rdr := csv.NewReader(bufio.NewReader(f))
	if _, err := rdr.Read(); err != nil {
		return Insights{}, err
	}

	// Setup channel & partials
	records := make(chan []string, ca.workers*2)
	var wg sync.WaitGroup

	// Structure for partial aggregation results
	type part struct {
		country map[struct{ C, P string }]struct {
			rev float64
			cnt int
		}
		prod   map[string]struct{ cnt, stock int }
		month  map[string]int
		region map[string]struct {
			rev  float64
			sold int
		}
	}
	partials := make([]part, ca.workers)

	// worker function for partial aggregation
	worker := func(idx int) {
		defer wg.Done()
		p := &partials[idx]

		// Initialize maps
		p.country = make(map[struct{ C, P string }]struct {
			rev float64
			cnt int
		})
		p.prod = make(map[string]struct{ cnt, stock int })
		p.month = make(map[string]int)
		p.region = make(map[string]struct {
			rev  float64
			sold int
		})

		// Process records
		for rec := range records {
			qty, _ := strconv.Atoi(rec[9])
			price, _ := strconv.ParseFloat(rec[8], 64)
			total := float64(qty) * price
			dt, _ := time.Parse("2006-01-02", rec[1])
			mon := dt.Format("2006-01")

			// Aggregate by country + product
			cp := struct{ C, P string }{rec[3], rec[6]}
			cv := p.country[cp]
			cv.rev += total
			cv.cnt++
			p.country[cp] = cv

			// Aggregate product purchases and stock quantity
			pv := p.prod[rec[6]]
			pv.cnt += qty
			st, _ := strconv.Atoi(rec[11])
			pv.stock = st
			p.prod[rec[6]] = pv

			// Aggregate monthly sales
			p.month[mon] += qty

			// Aggregate regional revenue and quantity sold
			rv := p.region[rec[4]]
			rv.rev += total
			rv.sold += qty
			p.region[rec[4]] = rv
		}
	}

	// start worker goroutines
	wg.Add(ca.workers)
	for i := 0; i < ca.workers; i++ {
		go worker(i)
	}

	// Feed records to workers concurrently
	go func() {
		defer close(records)
		for {
			rec, err := rdr.Read()
			if err == io.EOF {
				return
			}
			if err != nil {
				continue
			}
			records <- rec
		}
	}()
	wg.Wait()

	// Combine all partial results into final maps
	countryMap := make(map[struct{ C, P string }]struct {
		rev float64
		cnt int
	})
	prodMap := make(map[string]struct{ cnt, stock int })
	monthMap := make(map[string]int)
	regionMap := make(map[string]struct {
		rev  float64
		sold int
	})
	for _, p := range partials {
		for k, v := range p.country {
			cv := countryMap[k]
			cv.rev += v.rev
			cv.cnt += v.cnt
			countryMap[k] = cv
		}
		for k, v := range p.prod {
			pv := prodMap[k]
			pv.cnt += v.cnt
			if v.stock > pv.stock {
				pv.stock = v.stock
			}
			prodMap[k] = pv
		}
		for k, v := range p.month {
			monthMap[k] += v
		}
		for k, v := range p.region {
			rv := regionMap[k]
			rv.rev += v.rev
			rv.sold += v.sold
			regionMap[k] = rv
		}
	}

	//// Convert combined maps into sorted slices
	// Sort country-product revenue
	cr := make([]models.CountryRevenue, 0, len(countryMap))
	for k, v := range countryMap {
		cr = append(cr, models.CountryRevenue{Country: k.C, ProductName: k.P, TotalRevenue: v.rev, TransactionCount: v.cnt})
	}
	sort.Slice(cr, func(i, j int) bool { return cr[i].TotalRevenue > cr[j].TotalRevenue })

	// Sort top products by purchase count
	tp := make([]models.ProductFrequency, 0, len(prodMap))
	for k, v := range prodMap {
		tp = append(tp, models.ProductFrequency{ProductName: k, PurchaseCount: v.cnt, StockQuantity: v.stock})
	}
	sort.Slice(tp, func(i, j int) bool {
		if tp[i].PurchaseCount == tp[j].PurchaseCount {
			return tp[i].ProductName > tp[j].ProductName
		}
		return tp[i].PurchaseCount > tp[j].PurchaseCount
	})
	if len(tp) > 20 {
		tp = tp[:20]
	}

	// Sort monthly sales
	ms := make([]models.MonthlySales, 0, len(monthMap))
	for k, v := range monthMap {
		ms = append(ms, models.MonthlySales{Month: k, SalesVolume: v})
	}
	sort.Slice(ms, func(i, j int) bool {
		ti, _ := time.Parse("2006-01", ms[i].Month)
		tj, _ := time.Parse("2006-01", ms[j].Month)
		return ti.Before(tj)
	})

	// Sort region revenue in descending order
	rr := make([]models.RegionRevenue, 0, len(regionMap))
	for k, v := range regionMap {
		rr = append(rr, models.RegionRevenue{Region: k, TotalRevenue: v.rev, ItemsSold: v.sold})
	}
	sort.Slice(rr, func(i, j int) bool { return rr[i].TotalRevenue > rr[j].TotalRevenue })
	if len(rr) > 30 {
		rr = rr[:30]
	}

	return Insights{cr, tp, ms, rr}, nil
}
