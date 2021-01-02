package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var workers int

type Query struct {
	Host      string
	StartTime string
	EndTime   string
}

func ProcessQueriesFile(filename string) ([]Query, error) {

	var queries []Query

	lines, err := ReadCsv(filename)
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

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "SELECT queries benchmark for TimescaleDB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("number of workers:", workers)

		queries, err := ProcessQueriesFile("queries_test.csv")

		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println(queries[1].Host, queries[1].StartTime)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	selectCmd.Flags().IntVarP(&workers, "workers", "w", 1, "number of workers")
	selectCmd.MarkFlagRequired("workers")
}
