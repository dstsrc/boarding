package transposition

import (
	"sort"
)

type transposition struct {
	positions []int
	total     int
}

func New() *transposition {
	return &transposition{}
}

// Next return next transposition and completion mark
func (t *transposition) Next() ([]int, bool) {
	pp, done := next(t.positions)
	t.positions = pp
	if done {
		return nil, true
	}
	return pp, false
}

// Init return init params
func (t *transposition) Init(pos []int) []int {
	pp, total := getInitParams(pos)
	t.positions = pp
	t.total = total
	return pp
}

func getInitParams(pp []int) ([]int, int) {
	sort.Slice(pp, func(i, j int) bool {
		return pp[i] < pp[j]
	})

	total := 1
	for i := 1; i <= len(pp); i++ {
		total *= i
	}

	return pp, total
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
