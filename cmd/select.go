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
		fmt.Println("Number of workers:", workers)

		queries, err := query.ProcessQueriesFile("query_params.csv")

		if err != nil {
			fmt.Println("error:", err)
		}

		//fmt.Println("No of supplied queries:", len(queries))

		elapsedList := query.PerformQueries(workers, queries)

		fmt.Println("No of processed queries:", len(elapsedList))
		//fmt.Println(elapsedList)
		fmt.Println("Minimum query time:   ", elapsedList.Min())
		fmt.Println("Maximum query time:   ", elapsedList.Max())
		fmt.Println("Mean query time:      ", elapsedList.Mean())
		fmt.Println("Average query time:   ", elapsedList.Average())
		fmt.Println("Total processing time:", elapsedList.Sum())
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	selectCmd.Flags().IntVarP(&workers, "workers", "w", 1, "number of workers")
	selectCmd.MarkFlagRequired("workers")
}
