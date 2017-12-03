package main

import (
	"fmt"
	"math"
)

type square struct {
	x, y int
}

func main() {
	var input int
	fmt.Scanln(&input)
	fmt.Printf("part 1: %v\n", manhattan(location(input)))
	fmt.Printf("part 2: %v\n", stress(input))
}

func stress(target int) int {
	grid := map[square]int{location(1): 1}
	for n := 2; ; n++ {
		s := location(n)
		grid[s] = sumAdjacent(s, grid)
		if grid[s] > target {
			return grid[s]
		}
	}
}

func sumAdjacent(s square, grid map[square]int) int {
	sum := grid[square{s.x + 1, s.y}]
	sum += grid[square{s.x, s.y + 1}]
	sum += grid[square{s.x - 1, s.y}]
	sum += grid[square{s.x, s.y - 1}]
	sum += grid[square{s.x + 1, s.y + 1}]
	sum += grid[square{s.x + 1, s.y - 1}]
	sum += grid[square{s.x - 1, s.y + 1}]
	sum += grid[square{s.x - 1, s.y - 1}]
	return sum
}

func location(n int) square {
	// Jump to the correct ring, using the fact that the odd squares lie on the
	// diagonal.
	k := int(math.Ceil((math.Sqrt(float64(n)) - 1) / 2))
	m := (2*k + 1) * (2*k + 1)
	s := square{k, -k}
	// Move backwards around the ring one side at a time until we hit the right
	// square.
	if m > n {
		jump := min(m-n, 2*k)
		s.x -= jump
		m -= jump
	}
	if m > n {
		jump := min(m-n, 2*k)
		s.y += jump
		m -= jump
	}
	if m > n {
		jump := min(m-n, 2*k)
		s.x += jump
		m -= jump
	}
	if m > n {
		jump := min(m-n, 2*k)
		s.y -= jump
		m -= jump
	}
	return s
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func manhattan(s square) int {
	return abs(s.x) + abs(s.y)
}
