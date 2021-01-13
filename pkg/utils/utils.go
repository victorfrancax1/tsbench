// Package utils provides helper functions that will be used by other
// tsbench packages.
package utils

import (
	"encoding/csv"
	"os"
)

// ReadCsv receives a string that represents the name of the target CSV
// file and parses it into a slice of string slices, representing each row.
// This might be used to convert the CSV to other Go structs later on.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
