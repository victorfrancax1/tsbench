package benchmark

import (
	"fmt"
	"sort"
	"time"
)

// QueryTimes represents a slice of query durations, which will be the result of the benchmarks
// and will be used in order to gather desired metrics, such as mean, sum, min, max, avg.
type QueryTimes []time.Duration

// Len returns the length of QueryTimes object. It's necessary due to sort.
func (qt QueryTimes) Len() int { return len(qt) }

// Swap swaps two QueryTimes objects. It's necessary due to sort.
func (qt QueryTimes) Swap(i, j int) { qt[i], qt[j] = qt[j], qt[i] }

// Less compares two QueryTimes objects and returns the smaller one. It's necessary due to sort.
func (qt QueryTimes) Less(i, j int) bool { return qt[i] < qt[j] }

func (qt QueryTimes) average() time.Duration {
	var sum time.Duration = 0
	for _, t := range qt {
		sum = sum + t
	}
	avg := float64(sum) / float64(len(qt))
	return time.Duration(avg)
}

func (qt QueryTimes) min() time.Duration {
	sort.Sort(qt)
	return qt[0]
}

func (qt QueryTimes) max() time.Duration {
	sort.Sort(qt)
	return qt[len(qt)-1]
}

func (qt QueryTimes) sum() time.Duration {
	var sum time.Duration = 0
	for _, t := range qt {
		sum = sum + t
	}
	return sum
}

func (qt QueryTimes) mean() time.Duration {
	sort.Sort(qt)
	mIndex := len(qt) / 2

	if len(qt)%2 != 0 {
		return qt[mIndex]
	}

	return (qt[mIndex-1] + qt[mIndex]) / 2
}

// PrettyPrint is responsible for printing the output of the benchmark.
func (qt QueryTimes) PrettyPrint() {
	fmt.Println("No of processed queries:", len(qt))
	fmt.Println("Minimum query time:   ", qt.min())
	fmt.Println("Maximum query time:   ", qt.max())
	fmt.Println("Mean query time:      ", qt.mean())
	fmt.Println("Average query time:   ", qt.average())
	fmt.Println("Total processing time:", qt.sum())
}
