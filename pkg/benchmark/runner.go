package benchmark

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

func PerformQueries(numWorkers int, queries []SelectQuery) QueryTimes {
	var durations QueryTimes
	numQueries := len(queries)

	jobsList := generateJobs(queries)
	numJobs := len(jobsList)
	jobs := make(chan []SelectQuery, numJobs)
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
