package benchmark

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"tsbench/pkg/utils"
)

type SelectBenchmark struct {
	QueriesFileName string
	NumWorkers      int
	TsdbConnection  TsdbConnection
}

type SelectQuery struct {
	Host      string
	StartTime string
	EndTime   string
}

func (q SelectQuery) Execute(conn *pgxpool.Pool) time.Duration {
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
		fmt.Println("Query failed:", err)
	}

	elapsed := time.Since(start)

	rows.Close()
	return elapsed
}

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

func (sb SelectBenchmark) RunBenchmark() error {

	fmt.Println("Number of workers:", sb.NumWorkers)

	jobList, numQueries, err := sb.ProcessQueriesFile()

	if err != nil {
		return err
	}

	// this probably should have error handling
	queryTimes := PerformQueries(sb.NumWorkers, numQueries, jobList, sb.TsdbConnection)

	queryTimes.PrettyPrint()

	return nil
}
