package counter

type counter struct{}

func New() *counter {
	return &counter{}
}

func (c *counter) Count(pp []int) int {
	cp := make([]int, len(pp))
	copy(cp, pp)
	return countSlice(cp)
}

func countSlice(pp []int) int {
	if len(pp) == 0 {
		return 0
	}
	done := false
	cnt := 0
	for {
		pp, done = iterate(pp)
		if done {
			break
		}
		cnt++
	}
	return cnt
}

func iterate(pp []int) ([]int, bool) {
	index := len(pp) - 1
	for i := len(pp) - 1; i >= 0; i-- {
		if pp[i] > 0 {
			index = i
			break
		}
		if i == 0 {
			return pp, true
		}
	}

	seatNum := pp[index]
	for seatNum > 0 {
		if index == -1 || seatNum == -1 {
			return pp, false
		}

		if pp[index] <= seatNum && pp[index] != 0 {
			seatNum = pp[index]
			pp[index] = 0
		}
		index--
		seatNum--
	}
	return pp, false
}
