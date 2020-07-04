package modeling

import (
	"math"
	"math/rand"
	"time"
)

const target = 0.00001

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Counter interface {
	Count(pp []int) int
}

type mdl struct {
	counter Counter
}

func New(c Counter) *mdl {
	return &mdl{counter: c}
}

func (m *mdl) GetAVGTime(n int) float64 {
	s := make([]int, n)
	for i := range s {
		s[i] = i + 1
	}

	results := make([]int, 0)
	oldX := 0.0
	for {
		s = anotherTransposition(s)
		r := m.counter.Count(s)
		results = append(results, r)
		if len(results) < 10000 {
			continue
		}

		x := getX(results)
		delta := x - oldX
		if delta < 0 {
			delta *= -1
		}
		if delta < target {
			return x
		}
		oldX = x

	}
}

func getX(s []int) float64 {
	if len(s) == 0 {
		return 0
	}
	sum := 0
	for _, i := range s {
		sum += i
	}
	return float64(sum) / float64(len(s))
}

func getD(rr []int, x float64) float64 {
	if len(rr) == 0 {
		return 0
	}
	sum := 0.0
	for _, r := range rr {
		sum += (float64(r) - x) * (float64(r) - x)
	}
	return math.Sqrt(sum / float64(len(rr)))
}

func anotherTransposition(s []int) []int {
	for i := range s {
		j := rand.Intn(len(s))
		s[i], s[j] = s[j], s[i]
	}
	return s
}
