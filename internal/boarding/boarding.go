package boarding

type Counter interface {
	Count(pp []int) int
}

type Transposition interface {
	Init([]int) []int
	Next() ([]int, bool)
}

type boarding struct {
	counter       Counter
	transposition Transposition
}

func New(t Transposition, c Counter) *boarding {
	return &boarding{
		counter:       c,
		transposition: t,
	}
}

func (b *boarding) GetAVGTime(n int) float64 {
	pp := make([]int, 0)
	for i := 1; i <= n; i++ {
		pp = append(pp, i)
	}

	states := 0
	time := 0

	pp = b.transposition.Init(pp)
	t := b.counter.Count(pp)
	states++
	time += t

	done := false
	for {
		pp, done = b.transposition.Next()
		if done {
			break
		}
		t := b.counter.Count(pp)
		states++
		time += t
	}

	return float64(time) / float64(states)
}
