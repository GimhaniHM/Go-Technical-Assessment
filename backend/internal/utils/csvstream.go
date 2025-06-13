package utils

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

// StreamCSV filters rows into jobs via channel.
func StreamCSV(path string, jobs chan<- []string, errs chan<- error) {
	defer close(jobs)

	f, err := os.Open(path)
	if err != nil {
		errs <- err
		return
	}
	defer f.Close()

	csvr := csv.NewReader(bufio.NewReader(f))
	if _, err := csvr.Read(); err != nil {
		errs <- err
		return
	}

	for {
		rec, err := csvr.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			errs <- err
			return
		}
		jobs <- rec
	}
}

// // opens the file and streams its CSV records and then returns a channel of records or an error.
// func StreamCSV(path string) (<-chan []string, <-chan error) {
// 	records := make(chan []string)
// 	errs := make(chan error, 1)

// 	go func() {
// 		defer close(records)
// 		defer close(errs)

// 		f, err := os.Open(path)
// 		if err != nil {
// 			errs <- err
// 			return
// 		}
// 		defer f.Close()

// 		reader := bufio.NewReader(f)
// 		csvr := csv.NewReader(reader)

// 		if _, err := csvr.Read(); err != nil {
// 			errs <- err
// 			return
// 		}

// 		for {
// 			rec, err := csvr.Read()
// 			if err == io.EOF {
// 				return
// 			}
// 			if err != nil {
// 				errs <- err
// 				return
// 			}
// 			records <- rec
// 		}
// 	}()

// 	return records, errs
// }
