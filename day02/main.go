package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum1, sum2 int
	for scanner.Scan() {
		row := make([]int, 0)
		for _, f := range strings.Fields(scanner.Text()) {
			n, _ := strconv.Atoi(f)
			row = append(row, n)
		}
		sum1 += rowRange(row)
		sum2 += rowDivision(row)
	}
	fmt.Printf("part 1: %v\n", sum1)
	fmt.Printf("part 1: %v\n", sum2)
}

func rowRange(row []int) int {
	var min, max int
	for i, n := range row {
		if i == 0 {
			min = n
			max = n
		} else if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}
	return max - min
}

func rowDivision(row []int) int {
	for i, n := range row {
		for j, m := range row {
			if i != j && n%m == 0 {
				return n / m
			}
		}
	}
	return 0
}
