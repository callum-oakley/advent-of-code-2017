package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var groups []map[int]bool
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		n, pipes := parse(line)
		groups = connect(groups, n, pipes)
	}
	var zeroGroup map[int]bool
	for _, group := range groups {
		if group[0] {
			zeroGroup = group
		}
	}
	fmt.Printf("part 1: %v\n", len(zeroGroup))
	fmt.Printf("part 2: %v\n", len(groups))
}

func parse(s string) (int, []int) {
	var pipes []int
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) || r == ','
	})
	n, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Fatal(err)
	}
	for i := 2; i < len(fields); i++ {
		pipe, err := strconv.Atoi(fields[i])
		if err != nil {
			log.Fatal(err)
		}
		pipes = append(pipes, pipe)
	}
	return n, pipes
}

func connect(groups []map[int]bool, n int, pipes []int) []map[int]bool {
	connected := []map[int]bool{map[int]bool{n: true}}
	for _, group := range groups {
		if containsAnyOf(group, pipes) {
			connected[0] = union(connected[0], group)
		} else {
			connected = append(connected, group)
		}
	}
	return connected
}

func containsAnyOf(group map[int]bool, pipes []int) bool {
	for _, pipe := range pipes {
		if group[pipe] {
			return true
		}
	}
	return false
}

func union(x map[int]bool, y map[int]bool) map[int]bool {
	z := map[int]bool{}
	for n := range x {
		z[n] = true
	}
	for n := range y {
		z[n] = true
	}
	return z
}
