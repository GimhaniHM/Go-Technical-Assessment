package utils

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestReadTransactions(t *testing.T) {
	// Prepare a dummy CSV content with a single valid transaction row
	content := `transaction_id,transaction_date,user_id,country,region,product_id,product_name,category,price,quantity,total_price,stock_quantity,added_date
T123,2025-06-14,U1,USA,NA,P1,Prod1,Cat1,10.5,2,21.0,100,2025-06-13
`
	// Create a temporary directory and file to store the dummy CSV
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "test.csv")

	// Write the CSV content to the file
	if err := ioutil.WriteFile(file, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	got, err := ReadTransactions(file)
	if err != nil {
		t.Fatalf("ReadTransactions error: %v", err)
	}

	// Check that exactly one record was returned
	if len(got) != 1 {
		t.Fatalf("got %d records; want 1", len(got))
	}

	// Validate the content of the transaction
	tr := got[0]
	if tr.TransactionID != "T123" ||
		tr.Country != "USA" ||
		tr.Region != "NA" ||
		tr.Quantity != 2 ||
		tr.TotalPrice != 21.0 {
		t.Errorf("ReadTransactions record = %+v; want TransactionID=T123, Country=USA, Region=NA, Quantity=2, TotalPrice=21.0", tr)
	}
}
