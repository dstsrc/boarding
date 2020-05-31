package main

import (
	"fmt"

	"github.com/dstsrc/boarding/internal/boarding"

	"github.com/dstsrc/boarding/internal/counter"
	"github.com/dstsrc/boarding/internal/counterll"
	"github.com/dstsrc/boarding/internal/transposition"
)

func main() {
	n := 10

	t1 := transposition.New()
	c1 := counterll.New()
	b1 := boarding.New(t1, c1)

	avg1 := b1.GetAVGTime(n)
	fmt.Println(avg1)

	t2 := transposition.New()
	c2 := counter.New()
	b2 := boarding.New(t2, c2)

	avg2 := b2.GetAVGTime(n)
	fmt.Println(avg2)
}
