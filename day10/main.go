package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const knotSize = 256

var prefix = []byte{17, 31, 73, 47, 23}

func main() {
	var input string
	fmt.Scanln(&input)
	fields := strings.Split(input, ",")
	numericLengths := make([]byte, len(fields))
	for i, f := range fields {
		length, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		numericLengths[i] = byte(length)
	}
	fmt.Printf("part 1: %v\n", simpleHash(numericLengths))
	fmt.Printf("part 2: %v\n", denseHash(append([]byte(input), prefix...), 64))
}

func simpleHash(lengths []byte) int {
	sparse := sparseHash(lengths, 1)
	return int(sparse[0]) * int(sparse[1])
}

func denseHash(lengths []byte, rounds int) string {
	sparse := sparseHash(lengths, rounds)
	var hash string
	for block := 0; block < knotSize/16; block++ {
		offset := block * 16
		acc := sparse[offset]
		for i := offset + 1; i < offset+16; i++ {
			acc ^= sparse[i]
		}
		hash += fmt.Sprintf("%2.2x", acc)
	}
	return hash
}

func sparseHash(lengths []byte, rounds int) []byte {
	knot := make([]byte, knotSize)
	for i := 0; i < knotSize; i++ {
		knot[i] = byte(i)
	}
	var i, skip int
	for round := 0; round < rounds; round++ {
		for _, length := range lengths {
			reverse(knot, i, i+int(length)-1)
			i += int(length) + skip
			skip++
		}
	}
	return knot
}

// Reverse the elements between i and j inclusive
func reverse(knot []byte, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		knot[i%knotSize], knot[j%knotSize] = knot[j%knotSize], knot[i%knotSize]
	}
}
