package benchmark

// Benchmark is an interface for running benchmarks for one specific query, against a
// TimescaleDB instance.
type Benchmark interface {
	RunBenchmark() error
}
