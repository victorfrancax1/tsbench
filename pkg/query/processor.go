package query

import (
	"encoding/csv"
	"os"
	"time"
)

func worker(id int, jobs <-chan []Query, results chan<- time.Duration) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)

		for query := range j {
			elapsed := DoSelectQuery(query)
			results <- elapsed
		}

		fmt.Println("worker", id, "finished job", j)
	}
}

func generateJobs(queries []Query) [][]Query {
	//turn into map? 
}

func PerformQueries(numWorkers int, queries []Query) []time.Duration {
	var durations []time.Duration
	numQueries := len(queries)

	jobs := make(chan []Query)
	results := make(chan time.Duration, numQueries)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// generate jobs based on queries and feed them to job channel

	close(jobs)

	for a := 1; a <= numQueries; a++ {
		durations = append(durations, <-results)
	}
}
