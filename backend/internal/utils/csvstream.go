package utils

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/GimhaniHM/backend/internal/models"
)

// ReadTransactions reads a CSV file and converts each row into a Transaction struct.
// Returns a slice of transactions or an error if the file cannot be read.
func ReadTransactions(path string) ([]models.Transaction, error) {

	// Open the CSV file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	// skip header
	if _, err := r.Read(); err != nil {
		return nil, err
	}

	var out []models.Transaction

	// Read and parse each line until end of file
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// Parse required fields from the record
		td, _ := time.Parse("2006-01-02", rec[1])
		price, _ := strconv.ParseFloat(rec[8], 64)
		qty, _ := strconv.Atoi(rec[9])
		tot := float64(qty) * price
		stock, _ := strconv.Atoi(rec[11])
		ad, _ := time.Parse("2006-01-02", rec[12])

		// Create a Transaction struct from the record
		out = append(out, models.Transaction{
			TransactionID:   rec[0],
			TransactionDate: td,
			UserID:          rec[2],
			Country:         rec[3],
			Region:          rec[4],
			ProductID:       rec[5],
			ProductName:     rec[6],
			Category:        rec[7],
			Price:           price,
			Quantity:        qty,
			TotalPrice:      tot,
			StockQuantity:   stock,
			AddedDate:       ad,
		})
	}

	// Return the final list of transactions
	return out, nil
}
