package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var count1, count2 int
	for scanner.Scan() {
		phrase := scanner.Text()
		if isValid1(phrase) {
			count1++
		}
		if isValid2(phrase) {
			count2++
		}
	}
	fmt.Printf("part 1: %v\n", count1)
	fmt.Printf("part 2: %v\n", count2)
}

func isValid1(phrase string) bool {
	seen := make(map[string]bool)
	for _, word := range strings.Fields(phrase) {
		if seen[word] {
			return false
		}
		seen[word] = true
	}
	return true
}

func isValid2(phrase string) bool {
	seen := make(map[string]bool)
	for _, word := range strings.Fields(phrase) {
		normal := normalize(word)
		if seen[normal] {
			return false
		}
		seen[normal] = true
	}
	return true
}

func normalize(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
