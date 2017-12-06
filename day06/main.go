package main

import "fmt"

func main() {
	var banks []int
	for {
		var blocks int
		n, err := fmt.Scan(&blocks)
		if n != 1 || err != nil {
			break
		}
		banks = append(banks, blocks)
	}
	steps, size := stepsToLoop(banks)
	fmt.Printf("part 1: %v\n", steps)
	fmt.Printf("part 2: %v\n", size)
}

func stepsToLoop(banks []int) (int, int) {
	// Store the step at which a configuration of banks was seen, then we can
	// establish if any configuration has been seen before, and how many steps
	// ago.
	seen := map[string]int{hash(banks): 0}
	for i := 1; ; i++ {
		redistribute(banks)
		h := hash(banks)
		if j, ok := seen[h]; ok {
			return i, i - j
		}
		seen[h] = i
	}
}

func hash(banks []int) string {
	return fmt.Sprintf("%v", banks)
}

func redistribute(banks []int) {
	i, blocks := findLargest(banks)
	banks[i] = 0
	for blocks > 0 {
		i++
		banks[i%len(banks)]++
		blocks--
	}
}

func findLargest(banks []int) (int, int) {
	i := 0
	max := banks[0]
	for j := 1; j < len(banks); j++ {
		if banks[j] > max {
			max = banks[j]
			i = j
		}
	}
	return i, max
}
