package benchmark

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func worker(id int, conn *pgxpool.Pool, jobs <-chan []SelectQuery, results chan<- time.Duration) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job")

		for _, query := range j {
			elapsed := query.Execute(conn)
			results <- elapsed
		}

		//fmt.Println("worker", id, "finished job")
	}
}

func PerformQueries(numWorkers int, numQueries int, jobList [][]SelectQuery, tc TsdbConnection) QueryTimes {
	var durations QueryTimes

	conn, err := tc.GetConnectionPool()

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	numJobs := len(jobList)
	jobs := make(chan []SelectQuery, numJobs)
	results := make(chan time.Duration, numQueries)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, conn, jobs, results)
	}

	for _, j := range jobList {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numQueries; a++ {
		durations = append(durations, <-results)
	}

	return durations
}
