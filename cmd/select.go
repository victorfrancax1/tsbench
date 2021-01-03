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

		queries, err := query.ProcessQueriesFile("queries_test.csv")

		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println(queries[1].Host, queries[1].StartTime)

		elapsed := query.DoSelectQuery(queries[1])
		fmt.Println("elapsed:", elapsed)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	selectCmd.Flags().IntVarP(&workers, "workers", "w", 1, "number of workers")
	selectCmd.MarkFlagRequired("workers")
}
