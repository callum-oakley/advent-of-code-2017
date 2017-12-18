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
	var a, b int
	fmt.Scanf("Generator A starts with %v", &a)
	fmt.Scanf("Generator B starts with %v", &b)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		chA := generate(a, factorA, alwaysTrue)
		chB := generate(b, factorB, alwaysTrue)
		count := 0
		for i := 0; i < sample1; i++ {
			if uint16(<-chA) == uint16(<-chB) {
				count++
			}
		}
		ch1 <- count
	}()
	go func() {
		chA := generate(a, factorA, conditionA)
		chB := generate(b, factorB, conditionB)
		count := 0
		for i := 0; i < sample2; i++ {
			if uint16(<-chA) == uint16(<-chB) {
				count++
			}
		}
		ch2 <- count
	}()
	fmt.Printf("part 1: %v\n", <-ch1)
	fmt.Printf("part 2: %v\n", <-ch2)
}

func generate(n, factor int, condition func(int) bool) chan int {
	ch := make(chan int, 128)
	go func() {
		for {
			n = n * factor % divisor
			if condition(n) {
				ch <- n
			}
		}
	}()
	return ch
}
