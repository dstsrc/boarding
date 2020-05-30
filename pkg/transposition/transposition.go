package transposition

import (
	"sort"
)

type transposition struct {
	positions []int
	step      int
	total     int
}

func NewFromSlice(pp []int) *transposition {
	sort.Slice(pp, func(i, j int) bool {
		return pp[i] < pp[j]
	})

	total := 1
	for i := 1; i <= len(pp); i++ {
		total *= i
	}

	return NewTransposition(pp, total)
}

func NewTransposition(pp []int, total int) *transposition {
	return &transposition{
		positions: pp,
		total:     total,
		step:      0,
	}
}

func next(ss []int) ([]int, bool) {
	for i := len(ss) - 1; i > 0; i-- {
		if ss[i] > ss[i-1] {
			revers(ss[i:])
			j := minPos(ss[i-1], ss[i:])
			ss[i+j], ss[i-1] = ss[i-1], ss[i+j]
			return ss, false
		}
	}
	return nil, true
}

func minPos(t int, ss []int) int {
	for i, s := range ss {
		if s > t {
			return i
		}
	}
	return 0
}

func revers(ss []int) []int {
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return ss
}
