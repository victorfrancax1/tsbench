package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var tsdbConnString string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "tsbench",
	Long: "TimescaleDB coding assignment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TimescaleDB benchmarking tool")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tsbench.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".tsbench" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tsbench")
	}

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		if _, ok := err.(viper.ConfigParseError); ok {
			os.Exit(1)
		}
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	//If a config file is specified, it doesnt make sense to read env variables
	if cfgFile == "" {
		viper.AutomaticEnv()
		viper.SetEnvPrefix("tsbench")
	}
	tsdbConnString = viper.GetString("tsdb_conn_string")
	if tsdbConnString == "" {
		fmt.Println("TSDB Connection string not found.")
		os.Exit(1)
	}
}
