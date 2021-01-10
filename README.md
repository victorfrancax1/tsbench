# tsbench

tsbench is  a commandline tool that can be used to benchmark SELECT query performance across multiple workers/clients against a TimescaleDB instance. The tool should take as its input a CSV file (whose format is specified below) and a flag to specify the number of concurrent workers. After processing all the queries specified by the parameters in the CSV file, the tool should output a summary with the following stats:

- \# of queries processed
- total processing time across all queries
- the minimum query time (for a single query)
- the median query time
- the average query time
- and the maximum query time.
