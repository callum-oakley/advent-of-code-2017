package main

import (
	"fmt"
	"strings"
)

// Cube coordinates
// https://www.redblobgames.com/grids/hexagons/#coordinates-cube
type hex struct {
	x, y, z int
}

func (a hex) add(b hex) hex {
	return hex{a.x + b.x, a.y + b.y, a.z + b.z}
}

func main() {
	var input string
	fmt.Scanln(&input)
	fields := strings.Split(input, ",")
	var h hex
	var maxd int
	for _, f := range fields {
		h = h.add(fieldToHex(f))
		if d := distance(h); d > maxd {
			maxd = d
		}
	}
	fmt.Printf("part 1: %v\n", distance(h))
	fmt.Printf("part 1: %v\n", maxd)
}

func fieldToHex(s string) hex {
	switch s {
	case "n":
		return hex{0, 1, -1}
	case "ne":
		return hex{1, 0, -1}
	case "se":
		return hex{1, -1, 0}
	case "s":
		return hex{0, -1, 1}
	case "sw":
		return hex{-1, 0, 1}
	case "nw":
		return hex{-1, 1, 0}
	default:
		return hex{}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func distance(a hex) int {
	return (abs(a.x) + abs(a.y) + abs(a.z)) / 2
}
