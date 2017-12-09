package main

import "fmt"

func main() {
	var stream string
	fmt.Scanln(&stream)
	var depth, score, garbageCount int
	var garbage bool
	for i := 0; i < len(stream); i++ {
		if garbage {
			switch stream[i] {
			case '!':
				i++
			case '>':
				garbage = false
			default:
				garbageCount++
			}
		} else {
			switch stream[i] {
			case '{':
				depth++
			case '<':
				garbage = true
			case '}':
				score += depth
				depth--
			}
		}
	}
	fmt.Printf("part 1: %v\n", score)
	fmt.Printf("part 2: %v\n", garbageCount)
}
