package benchmark

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// worker function represents the task that each concurrent goroutine will be responsible for, which is
// receiving a list of queries, executing them, and sending the durations to the results channel.
// If an error occurs, it will be forwarded to a separate channel instead.
func worker(id int, conn *pgxpool.Pool, jobs <-chan []SelectQuery, results chan<- time.Duration, errors chan<- error) {
	for j := range jobs {
		for _, query := range j {
			elapsed, err := query.Execute(conn)
			if err != nil {
				errors <- err
			}
			results <- elapsed
		}
	}
}

// PerformQueries is responsible for orchestrating the concurrent workers, as well as handling input and output channels
// It will trigger workers for each SelectQuery slice in jobList, as well as collect results (and eventual errors).
func PerformQueries(numWorkers int, numQueries int, jobList [][]SelectQuery, tc TsdbConnection) (QueryTimes, error) {
	var durations QueryTimes

	conn, err := tc.GetConnectionPool()

	if err != nil {
		return durations, err
	}

	defer conn.Close()

	numJobs := len(jobList)
	jobs := make(chan []SelectQuery, numJobs)
	results := make(chan time.Duration, numQueries)
	errors := make(chan error, 0)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, conn, jobs, results, errors)
	}

	for _, j := range jobList {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numQueries; a++ {
		select {
		case r := <-results:
			durations = append(durations, r)
		case err := <-errors:
			return durations, err
		}
	}
	return durations, nil
}
