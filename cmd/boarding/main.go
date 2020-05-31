package main

import (
	"fmt"

	"github.com/dstsrc/boarding/internal/counter"

	"github.com/dstsrc/boarding/internal/boarding"

	"github.com/dstsrc/boarding/internal/transposition"
)

func main() {
	t := transposition.New()
	c := counter.New()
	b := boarding.New(t, c)

	avg := b.GetAVGTime(10)
	fmt.Print(avg)
}
