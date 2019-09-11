package algorithm

import (
	"fmt"
	"testing"
)

func TestIntervalMerge(t *testing.T) {
	var intervals = []Interval{
		Interval{1, 2}, Interval{7, 8}, Interval{8, 10}, Interval{2, 3}, Interval{4, 5}, }

	var res = []Interval{
		Interval{1, 3}, Interval{4, 5}, Interval{7, 10},}

	res1 := IntervalMerge(intervals)

	fmt.Println(IntervalsEqual(res, res1))
}
