package counterll

import "fmt"

type node struct {
	next  *node
	value int
}

type counter struct{}

func New() *counter {
	return &counter{}
}

// Count return the number of iterations for which everyone is sitting
func (c *counter) Count(pp []int) int {
	return countSlice(pp)
}

func countSlice(pp []int) int {
	n := toLL(pp)
	return countLL(n)
}

func countLL(n *node) int {
	cnt := 0
	for n != nil {
		n = iterate(n)
		cnt++
	}
	return cnt
}

// arrToArr for simplify the test
func arrToArr(pp []int) []int {
	return fromLL(iterate(toLL(pp)))
}

func iterate(head *node) *node {
	if head == nil {
		return nil
	}
	seatNum := head.value
	current := head
	next := head.next
	var prev *node = nil

	set := false
	for current != nil {
		if prev != nil && !set {
			head = prev
			set = true
		}

		if current.value <= seatNum {
			seatNum = current.value
			if prev != nil {
				prev.next = next
			}
			current = next
			if current != nil {
				next = current.next
			}
			if !set {
				head = current
			}
		} else {
			prev = current
			current = current.next
			if current != nil {
				next = current.next
			}
		}

		seatNum--
		if seatNum < 1 {
			break
		}
	}

	return head
}

func pprint(n *node) {
	for n != nil {
		fmt.Print(n.value)
		n = n.next
	}
}

func fromLL(n *node) []int {
	s := make([]int, 0)
	for n != nil {
		s = append(s, n.value)
		n = n.next
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func toLL(pp []int) *node {
	if len(pp) == 0 {
		return nil
	}

	n := &node{value: pp[len(pp)-1]}
	head := n

	for i := len(pp) - 1; i > 0; i-- {
		nextNode := &node{value: pp[i-1]}
		n.next = nextNode
		n = nextNode
	}
	return head
}
