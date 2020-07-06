package counter

type counter struct{}

func New() *counter {
	return &counter{}
}

// Count return the number of iterations for which everyone is sitting
func (c *counter) Count(pp []int) int {
	cp := make([]int, len(pp))
	copy(cp, pp)
	return countSliceV2(cp)
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

func countSliceV2(pp []int) int {
	if len(pp) == 0 {
		return 0
	}
	done := false
	cnt := 0
	for {
		pp, done = iterateV2(pp)
		cnt++
		if done {
			break
		}
	}
	return cnt
}

func iterateV2(pp []int) ([]int, bool) {
	if len(pp) == 0 {
		return []int{}, true
	}
	index := len(pp) - 1
	seatNum := pp[index]
	stop := index + 1
	setStop := false
	for {
		if index == -1 || seatNum == -1 {
			return pp[:stop], len(pp[:stop]) == 0
		}
		if pp[index] == 0 {
			index--
			if !setStop {
				stop--
			}
			continue
		}

		seat := false
		if pp[index] <= seatNum {
			seatNum = pp[index]
			pp[index] = 0
			seat = true
		}
		seatNum--
		if !setStop && seat {
			stop--
		}

		if !seat && !setStop {
			setStop = true
		}
		index--
	}
}
