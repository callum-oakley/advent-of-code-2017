package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)
	var sum1, sum2 int
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+1)%len(input)] {
			sum1 += int(input[i] - '0')
		}
		if input[i] == input[(i+(len(input)/2))%len(input)] {
			sum2 += int(input[i] - '0')
		}
	}
	fmt.Printf("part 1: %v\n", sum1)
	fmt.Printf("part 1: %v\n", sum2)
}
