package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"tsbench/pkg/query"
)

var workers int

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "SELECT queries benchmark for TimescaleDB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("number of workers:", workers)

		queries, err := query.ProcessQueriesFile("query_params.csv")

		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println("No of supplied queries:", len(queries))

		elapsedList := query.PerformQueries(workers, queries)

		fmt.Println("No of processed queries:", len(elapsedList))
		fmt.Println(elapsedList)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	selectCmd.Flags().IntVarP(&workers, "workers", "w", 1, "number of workers")
	selectCmd.MarkFlagRequired("workers")
}
