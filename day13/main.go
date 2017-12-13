package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	// Just store depth to range as given
	firewall := map[int]int{}
	var totalDepth int
	for {
		var depth, rng int
		_, err := fmt.Scanf("%v: %v", &depth, &rng)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		firewall[depth] = rng
		totalDepth = depth + 1
	}
	fmt.Printf("part 1: %v\n", severityOfTrip(firewall, totalDepth))
	fmt.Printf("part 1: %v\n", shortestSafeDelay(firewall, totalDepth))
}

func severityOfTrip(firewall map[int]int, totalDepth int) int {
	var severity int
	for depth := 0; depth < totalDepth; depth++ {
		// When the packet reaches a geven depth, the scanner will have
		// advanced depth units, and has a periodicity of (range - 1) * 2, so
		// we needn't advance a simulated scanner step by step.
		if rng := firewall[depth]; rng != 0 && depth%((rng-1)*2) == 0 {
			severity += depth * rng
		}
	}
	return severity
}

func shortestSafeDelay(firewall map[int]int, totalDepth int) int {
	for delay := 0; ; delay++ {
		if isSafe(firewall, totalDepth, delay) {
			return delay
		}
	}
}

func isSafe(firewall map[int]int, totalDepth int, delay int) bool {
	for depth := 0; depth < totalDepth; depth++ {
		if rng := firewall[depth]; rng != 0 && (delay+depth)%((rng-1)*2) == 0 {
			return false
		}
	}
	return true
}
