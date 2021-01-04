package query

import (
	"sort"
	"time"
)

type QueryTimes []time.Duration

func (qt QueryTimes) Len() int { return len(qt) }

func (qt QueryTimes) Swap(i, j int) { qt[i], qt[j] = qt[j], qt[i] }

func (qt QueryTimes) Less(i, j int) bool { return qt[i] < qt[j] }

func (qt QueryTimes) Average() time.Duration {
	var sum time.Duration = 0
	for _, t := range qt {
		sum = sum + t
	}
	avg := float64(sum) / float64(len(qt))
	return time.Duration(avg)
}

func (qt QueryTimes) Min() time.Duration {
	sort.Sort(qt)
	return qt[0]
}

func (qt QueryTimes) Max() time.Duration {
	sort.Sort(qt)
	return qt[len(qt)-1]
}

func (qt QueryTimes) Sum() time.Duration {
	var sum time.Duration = 0
	for _, t := range qt {
		sum = sum + t
	}
	return sum
}

func (qt QueryTimes) Mean() time.Duration {
	sort.Sort(qt)
	mIndex := len(qt) / 2

	if len(qt)%2 != 0 {
		return qt[mIndex]
	}

	return (qt[mIndex-1] + qt[mIndex]) / 2
}

func worker(id int, jobs <-chan []Query, results chan<- time.Duration) {
	for j := range jobs {
		//fmt.Println("worker", id, "started  job")

		for _, query := range j {
			elapsed := DoSelectQuery(query)
			results <- elapsed
		}

		//fmt.Println("worker", id, "finished job")
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

func PerformQueries(numWorkers int, queries []Query) QueryTimes {
	var durations QueryTimes
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
