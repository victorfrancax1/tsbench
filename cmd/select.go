package cmd

import (
	"github.com/spf13/cobra"
	"tsbench/pkg/benchmark"
)

var numWorkers int

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "SELECT queries benchmark for TimescaleDB",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tsdbConnection := benchmark.TsdbConnection{
			TsdbConnString: tsdbConnString,
		}

		sb := benchmark.SelectBenchmark{
			QueriesFileName: args[0],
			NumWorkers:      numWorkers,
			TsdbConnection:  tsdbConnection,
		}

		if err := sb.RunBenchmark(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	selectCmd.Flags().IntVarP(&numWorkers, "workers", "w", 1, "number of workers")
	selectCmd.MarkFlagRequired("workers")
}
