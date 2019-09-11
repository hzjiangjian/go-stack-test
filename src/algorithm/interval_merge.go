package algorithm

import "sort"

type Interval struct {
	Start int
	End   int
}

func (in *Interval) Merge(dst Interval) bool {
	if !in.Contains(dst.Start) && !dst.Contains(in.Start) {
		return false
	}

	if in.Start > dst.Start {
		in.Start = dst.Start
	}

	if in.End < dst.End {
		in.End = dst.End
	}

	return true
}

func (in *Interval) Contains(i int) bool {
	return i >= in.Start && i <= in.End
}

func IntervalsEqual(src, dst []Interval) bool {
	if len(src) != len(dst) {
		return false
	}

	for i := 0; i < len(src); i++ {
		if src[i].Start != dst[i].Start || src[i].End != dst[i].End {
			return false
		}
	}

	return true
}

type IntervalSlice []Interval

func (is IntervalSlice) Len() int { return len(is) }

func (is IntervalSlice) Swap(i, j int) { is[i], is[j] = is[j], is[i] }

func (is IntervalSlice) Less(i, j int) bool { return is[i].Start < is[j].Start }

func IntervalMerge(intervals []Interval) (res []Interval) {
	res = make([]Interval, 0)

	is := IntervalSlice(intervals)
	sort.Sort(is)

	i, j := 0, 0

	for ; j < is.Len(); {
		if i >= len(res) {
			res = append(res, Interval{})
			res[i] = is[j]
			j++
		} else {
			if res[i].Merge(is[j]) {
				j++
			} else {
				i++
			}
		}
	}

	return res
}
