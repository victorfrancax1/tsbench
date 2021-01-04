package query

import (
	"encoding/csv"
	"os"
)

type Query struct {
	Host      string
	StartTime string
	EndTime   string
}

func ProcessQueriesFile(filename string) ([]Query, error) {

	var queries []Query

	lines, err := readCsv(filename)
	if err != nil {
		return []Query{}, err
	}

	for _, line := range lines {
		query := Query{
			Host:      line[0],
			StartTime: line[1],
			EndTime:   line[2],
		}
		queries = append(queries, query)
	}

	return queries, nil
}

func readCsv(filename string) ([][]string, error) {

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
