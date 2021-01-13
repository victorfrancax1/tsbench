// tsbench is a command-line interface that aims to benchmark the performance
// of queries across multiple workers/clients against a TimescaleDB instance.
package main

import "github.com/victorfrancax1/tsbench/cmd"

func main() {
	cmd.Execute()
}
