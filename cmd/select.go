package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var workers int

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "SELECT queries benchmark for TimescaleDB",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("number of workers:", workers)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	selectCmd.Flags().IntVarP(&workers, "workers", "w", 1, "number of workers")
	selectCmd.MarkFlagRequired("workers")
}
