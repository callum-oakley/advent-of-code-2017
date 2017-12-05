package main

import "fmt"

func main() {
	var maze []int
	for {
		var instruction int
		n, err := fmt.Scan(&instruction)
		if n != 1 || err != nil {
			break
		}
		maze = append(maze, instruction)
	}
	fmt.Printf("part 1: %v\n", steps(
		func(jmp int) int { return jmp + 1 },
		maze,
	))
	fmt.Printf("part 2: %v\n", steps(
		func(jmp int) int {
			if jmp >= 3 {
				return jmp - 1
			}
			return jmp + 1
		},
		maze,
	))
}

func steps(transformJmp func(int) int, maze []int) int {
	// Take a copy of maze so we can mutate it without changing the original.
	mazeCpy := make([]int, len(maze))
	copy(mazeCpy, maze)
	var steps int
	for i := 0; i >= 0 && i < len(mazeCpy); steps++ {
		jmp := mazeCpy[i]
		mazeCpy[i] = transformJmp(mazeCpy[i])
		i += jmp
	}
	return steps
}
