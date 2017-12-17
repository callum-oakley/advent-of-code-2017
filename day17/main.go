package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var steps int
	fmt.Scan(&steps)
	fmt.Printf("\rpart 1: %v\n", spin(steps, 2018).Next().Value)
	fmt.Printf("\rpart 2: %v\n", focussedSpin(steps, 50000001))
}

func spin(steps, maxn int) *ring.Ring {
	pos := &ring.Ring{}
	for n := 1; n < maxn; n++ {
		r := &ring.Ring{Value: n}
		pos.Move(steps).Link(r)
		pos = r
	}
	return pos
}

func focussedSpin(steps, maxn int) int {
	// If we index the values from 0, then the only element we care about is
	// always at index 1, so we can ignore updates anywhere else in the ring.
	var valueAfterZero, pos int
	for n := 1; n < maxn; n++ {
		pos = ((pos + steps) % n) + 1
		if pos == 1 {
			valueAfterZero = n
		}
	}
	return valueAfterZero
}
