package benchmark

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func worker(id int, conn *pgxpool.Pool, jobs <-chan []SelectQuery, results chan<- time.Duration, errors chan<- error) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job")

		for _, query := range j {
			elapsed, err := query.Execute(conn)
			if err != nil {
				errors <- err
			}
			results <- elapsed
		}

		//fmt.Println("worker", id, "finished job")
	}
}

func PerformQueries(numWorkers int, numQueries int, jobList [][]SelectQuery, tc TsdbConnection) (QueryTimes, error) {
	var durations QueryTimes

	conn, err := tc.GetConnectionPool()

	if err != nil {
		fmt.Println(err)
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
