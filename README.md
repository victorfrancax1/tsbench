# tsbench

tsbench is a command-line tool that can be used to benchmark SELECT query performance across multiple workers/clients against a TimescaleDB instance. The tool should take as its input a CSV file (whose format is specified below) and a flag to specify the number of concurrent workers. After processing all the queries specified by the parameters in the CSV file, the tool should output a summary with the following stats:

- \# of queries processed
- total processing time across all queries
- the minimum query time (for a single query)
- the median query time
- the average query time
- and the maximum query time.

## Instalation

```bash
# Fetch tsbench and its dependencies
go get github.com/victorfrancax1/tsbench
cd $GOPATH/src/github.com/timescale/tsbench
go get ./...

# Install the binary
cd $GOPATH/src/github.com/victorfrancax1/tsbench
go install
```

## Configuration
Tsbench supports configs either through a config file, or through ENV variables. It will look first for a `.tsbench.yaml` file in $HOME, and if that fails, it will look for ENV variables that start with the prefix `TSBENCH`. It is also possible to specify a config file using `--config` flag.

Note: If a config file is specified with `--config`, tsbench won't look for ENV variables.

- Example config file:
```yaml
tsdb_conn_string: "postgresql://username@127.0.0.1/database"
```
- Equivalent ENV variable:
```bash
TSBENCH_TSDB_CONN_STRING="postgresql://username@127.0.0.1/database"
```


## Usage
Tsbench has only one command, `select`, which will be responsible for benchmarking our SELECT queries. To check out its options, run `tsbench select --help`.

```
Query benchmarking tool for TimescaleDB

Usage:
  tsbench [command]

Available Commands:
  help        Help about any command
  select      SELECT queries benchmark for TimescaleDB

Flags:
      --config string   config file (default is $HOME/.tsbench.yaml)
  -h, --help            help for tsbench

Use "tsbench [command] --help" for more information about a command.
```
### Examples
```
tsbench select --workers 10 query_params.csv 
```
```
tsbench select -w 5 query_params.csv --config .myconfig.yaml
```
