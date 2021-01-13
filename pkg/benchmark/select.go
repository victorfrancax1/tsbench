// Package benchmark is responsible for performing benchmarks of queries against TimescaleDB
// instances. In tsbench, it is called by package cmd, but it can be used standalone as well.

package benchmark

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"tsbench/pkg/utils"
)

// SelectBenchmark represents the benchmark that will be made to SELECT queries,
// and holds parameters as the input file name, the number of concurrent workers that will be used,
// and a TsdbConnection.
type SelectBenchmark struct {
	QueriesFileName string
	NumWorkers      int
	TsdbConnection  TsdbConnection
}

// SelectQuery represents the SELECT query that will be benchmarked against a
// TimescaleDB instance.
type SelectQuery struct {
	Host      string
	StartTime string
	EndTime   string
}

// Execute will run the respective query against TimescaleDB, given a connection
// (pool), returning the duration of the query.
func (q SelectQuery) Execute(conn *pgxpool.Pool) (time.Duration, error) {
	var elapsed time.Duration
	queryString := fmt.Sprintf(`
		SELECT time_bucket('1 minute', ts) as one_min,
			MAX(usage) as max_usage,
			MIN(usage) as min_usage
		FROM cpu_usage
		WHERE host = '%s'
			AND ts >= '%s'
			AND ts <= '%s'
		GROUP BY one_min`,
		q.Host,
		q.StartTime,
		q.EndTime)

	start := time.Now()

	rows, err := conn.Query(context.Background(), queryString)
	if err != nil {
		return elapsed, err
	}

	elapsed = time.Since(start)

	rows.Close()
	return elapsed, nil
}

// ProcessQueriesFile is responsible for transforming the parsed CSV input file,
// converting it to SelectQuery structs and then using generateJobs to make the queries
// ready for the job executions.
func (sb SelectBenchmark) ProcessQueriesFile() ([][]SelectQuery, int, error) {

	var queries []SelectQuery

	lines, err := utils.ReadCsv(sb.QueriesFileName)
	if err != nil {
		return [][]SelectQuery{}, 0, err
	}

	for i, line := range lines {
		if i == 0 {
			// skip header
			continue
		}
		query := SelectQuery{
			Host:      line[0],
			StartTime: line[1],
			EndTime:   line[2],
		}
		queries = append(queries, query)
	}

	numQueries := len(queries)
	jobList := generateJobs(queries)

	return jobList, numQueries, nil
}

// generateJobs is a helper function that takes in a slice of SelectQuery, and
// groups queries that are meant for the same host (benchmark constraint).
func generateJobs(queries []SelectQuery) [][]SelectQuery {
	var jobsMap = make(map[string][]SelectQuery)

	for _, q := range queries {
		jobsMap[q.Host] = append(jobsMap[q.Host], q)
	}

	var jobs [][]SelectQuery
	for _, j := range jobsMap {
		jobs = append(jobs, j)
	}
	return jobs
}

// RunBenchmark is responsible for wrapping of the steps required by the benchmark.
// It will be called by tsbench/cmd.
func (sb SelectBenchmark) RunBenchmark() error {

	fmt.Println("Number of workers:", sb.NumWorkers)

	jobList, numQueries, err := sb.ProcessQueriesFile()

	if err != nil {
		return err
	}

	queryTimes, err := PerformQueries(sb.NumWorkers, numQueries, jobList, sb.TsdbConnection)

	if err != nil {
		return err
	}

	queryTimes.PrettyPrint()

	return nil
}
