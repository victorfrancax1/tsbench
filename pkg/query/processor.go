package query

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan []Query, results chan<- time.Duration) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)

		for _, query := range j {
			elapsed := DoSelectQuery(query)
			results <- elapsed
		}

		fmt.Println("worker", id, "finished job", j)
	}
}

func generateJobs(queries []Query) [][]Query {
	var jobsMap = make(map[string][]Query)

	for _, q := range queries {
		jobsMap[q.Host] = append(jobsMap[q.Host], q)
	}

	var jobs [][]Query
	for _, j := range jobsMap {
		jobs = append(jobs, j)
	}
	return jobs
}

func PerformQueries(numWorkers int, queries []Query) []time.Duration {
	var durations []time.Duration
	numQueries := len(queries)

	jobsList := generateJobs(queries)
	numJobs := len(jobsList)
	jobs := make(chan []Query, numJobs)
	results := make(chan time.Duration, numQueries)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for _, j := range jobsList {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numQueries; a++ {
		durations = append(durations, <-results)
	}

	return durations
}
