package main

import "fmt"

const (
	sample1 = 4e7
	sample2 = 5e6
	divisor = 2147483647
	factorA = 16807
	factorB = 48271
)

var (
	conditionA = func(a int) bool { return a%4 == 0 }
	conditionB = func(b int) bool { return b%8 == 0 }
	alwaysTrue = func(int) bool { return true }
)

func main() {
	var startA, startB, count1, count2, a, b int
	fmt.Scanf("Generator A starts with %v", &startA)
	fmt.Scanf("Generator B starts with %v", &startB)
	a, b = startA, startB
	for i := 0; i < sample1; i++ {
		a, b = generate(a, factorA, alwaysTrue), generate(b, factorB, alwaysTrue)
		if uint16(a) == uint16(b) {
			count1++
		}
	}
	a, b = startA, startB
	for i := 0; i < sample2; i++ {
		a, b = generate(a, factorA, conditionA), generate(b, factorB, conditionB)
		if uint16(a) == uint16(b) {
			count2++
		}
	}
	fmt.Printf("part 1: %v\n", count1)
	fmt.Printf("part 1: %v\n", count2)
}

func generate(n, factor int, condition func(int) bool) int {
	n = n * factor % divisor
	for !condition(n) {
		n = n * factor % divisor
	}
	return n
}
