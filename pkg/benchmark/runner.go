package benchmark

import (
	"time"
)

func worker(id int, jobs <-chan []SelectQuery, results chan<- time.Duration) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job")

		for _, query := range j {
			elapsed := query.Execute()
			results <- elapsed
		}

		//fmt.Println("worker", id, "finished job")
	}
}

func PerformQueries(numWorkers int, numQueries int, jobList [][]SelectQuery) QueryTimes {
	var durations QueryTimes

	numJobs := len(jobList)
	jobs := make(chan []SelectQuery, numJobs)
	results := make(chan time.Duration, numQueries)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
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
