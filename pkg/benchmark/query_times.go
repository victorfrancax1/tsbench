package benchmark

import (
	"fmt"
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

func (qt QueryTimes) PrettyPrint() {
	fmt.Println("No of processed queries:", len(qt))
	fmt.Println("Minimum query time:   ", qt.Min())
	fmt.Println("Maximum query time:   ", qt.Max())
	fmt.Println("Mean query time:      ", qt.Mean())
	fmt.Println("Average query time:   ", qt.Average())
	fmt.Println("Total processing time:", qt.Sum())
}
