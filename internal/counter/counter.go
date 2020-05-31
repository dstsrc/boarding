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
			return pp[:0], true
		}
	}
	stop := index + 1
	seatNum := pp[index]
	for seatNum > 0 {
		if index == -1 || seatNum == -1 {
			return pp[:stop], false
		}
		if pp[index] != 0 {
			if pp[index] <= seatNum {
				seatNum = pp[index]
				pp[index] = 0
			}
			seatNum--
		}
		index--
	}
	return pp[:stop], false
}
