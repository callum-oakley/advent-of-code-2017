package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type disc struct {
	weight  int
	holding []string
}

func main() {
	discs := map[string]disc{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		n, d := parse(line)
		discs[n] = d
	}
	root, err := findRoot(discs)
	if err != nil {
		log.Fatal(err)
	}
	_, c := findTotalOrBalance(discs, root)
	fmt.Printf("part 1: %v\n", root)
	fmt.Printf("part 1: %v\n", c)
}

func findTotalOrBalance(discs map[string]disc, node string) (int, int) {
	total := discs[node].weight
	subtotals := map[string]int{}
	for _, n := range discs[node].holding {
		w, c := findTotalOrBalance(discs, n)
		if c != 0 {
			return 0, c
		}
		total += w
		subtotals[n] = w
	}
	if !allEqual(subtotals) {
		return 0, correction(discs, subtotals)
	}
	return total, 0
}

func correction(discs map[string]disc, subtotals map[string]int) int {
	targetWeight := mostCommonWeight(subtotals)
	strayNode, strayWeight, err := uniqueWeight(subtotals)
	if err != nil {
		log.Fatal(err)
	}
	return discs[strayNode].weight - (strayWeight - targetWeight)
}

func allEqual(m map[string]int) bool {
	return len(frequencies(m)) <= 1
}

func mostCommonWeight(weights map[string]int) int {
	var maxWeight, maxFreq int
	for weight, freq := range frequencies(weights) {
		if freq > maxFreq {
			maxFreq = freq
			maxWeight = weight
		}
	}
	return maxWeight
}

func uniqueWeight(weights map[string]int) (string, int, error) {
	// Assumes throughout that weights are positive
	var weight int
	for w, freq := range frequencies(weights) {
		if freq == 1 {
			weight = w
			break
		}
	}
	for node, w := range weights {
		if w == weight {
			return node, weight, nil
		}
	}
	return "", 0, fmt.Errorf("Couldn't find a unique weight (target: %v)!\n", weight)
}

func frequencies(m map[string]int) map[int]int {
	freqs := map[int]int{}
	for _, v := range m {
		freqs[v]++
	}
	return freqs
}

func findRoot(discs map[string]disc) (string, error) {
	children := map[string]bool{}
	for _, d := range discs {
		for _, child := range d.holding {
			children[child] = true
		}
	}
	for name := range discs {
		if !children[name] {
			return name, nil
		}
	}
	return "", errors.New("Couldn't find root!")
}

func parse(s string) (string, disc) {
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return unicode.IsSpace(r) || r == ',' || r == '(' || r == ')'
	})
	w, _ := strconv.Atoi(fields[1])
	d := disc{weight: w}
	for i := 3; i < len(fields); i++ {
		d.holding = append(d.holding, fields[i])
	}
	return fields[0], d
}
