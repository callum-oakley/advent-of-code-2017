package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	var moves []func(string) string
	for _, field := range strings.Split(input, ",") {
		moves = append(moves, parse(field))
	}
	dancers := "abcdefghijklmnop"
	fmt.Printf("part 1: %v\n", dance(moves, dancers))
	fmt.Printf("part 2: %v\n", danceMany(1e9, moves, dancers))
}

func danceMany(n int, moves []func(string) string, dancers string) string {
	// The dance repeats every findPeriod(moves, dancers) dances, so we need
	// only repeat the dance n%findPeriod(moves, dancers) times for the final
	// arrangement.
	for i := 0; i < n%findPeriod(moves, dancers); i++ {
		dancers = dance(moves, dancers)
	}
	return dancers
}

func findPeriod(moves []func(string) string, dancers string) int {
	initial := dancers
	for i := 0; ; i++ {
		dancers = dance(moves, dancers)
		if dancers == initial {
			return i + 1
		}
	}
}

func dance(moves []func(string) string, dancers string) string {
	for _, move := range moves {
		dancers = move(dancers)
	}
	return dancers
}

func parse(field string) func(string) string {
	switch field[0] {
	case 's':
		var x int
		fmt.Sscanf(field, "s%v", &x)
		return func(d string) string {
			return d[len(d)-x:] + d[:len(d)-x]
		}
	case 'x':
		var i, j int
		fmt.Sscanf(field, "x%v/%v", &i, &j)
		return func(d string) string {
			return swap(i, j, d)
		}
	case 'p':
		var a, b string
		fmt.Sscanf(field, "p%1v/%1v", &a, &b)
		return func(d string) string {
			return swap(strings.Index(d, a), strings.Index(d, b), d)
		}
	}
	return func(d string) string { return d }
}

func swap(i, j int, s string) string {
	if i > j {
		i, j = j, i
	}
	return s[:i] + s[j:j+1] + s[i+1:j] + s[i:i+1] + s[j+1:]
}
