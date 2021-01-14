// Package cmd is responsible to handle all of the command-line interface logic of
// the application. It is built with sf13/cobra and spf12/viper.
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/victorfrancax1/tsbench/pkg/utils"
)

var cfgFile string
var tsdbConnString string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "tsbench",
	Long: "Query benchmarking tool for TimescaleDB",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		errMessage := fmt.Sprintf("[ERROR] %v\n", err)
		fmt.Fprintln(os.Stderr, errMessage)
		os.Exit(1)
	}
}

// init is used here to initialize config (file or ENV) settings and flags
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tsbench.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		utils.HandleFailure(err)

		// Search config in home directory with name ".tsbench" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tsbench")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok || cfgFile != "" {
			utils.HandleFailure(err)
		}
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// If a config file is specified, it doesnt make sense to read ENV variables.
	if cfgFile == "" {
		viper.AutomaticEnv()
		viper.SetEnvPrefix("tsbench")
	}
	tsdbConnString = viper.GetString("tsdb_conn_string")
	if tsdbConnString == "" {
		err := errors.New("TimescaleDB connection string not found")
		utils.HandleFailure(err)
	}
}
