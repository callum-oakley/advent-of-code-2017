package main

import (
	"fmt"
	"strconv"
)

const knotSize = 256
const gridSize = 128

var prefix = []byte{17, 31, 73, 47, 23}

type pt struct {
	x, y int
}

type region map[pt]bool

func main() {
	var input string
	fmt.Scanln(&input)
	var usedCount int
	used := map[pt]bool{}
	for y := 0; y < gridSize; y++ {
		hash := denseHash(
			append([]byte(fmt.Sprintf("%v-%v", input, y)), prefix...),
			64,
		)
		for x, d := range hash {
			if d == '1' {
				used[pt{x, y}] = true
				usedCount++
			}
		}
	}
	fmt.Printf("part 1: %v\n", usedCount)
	fmt.Printf("part 2: %v\n", countRegions(used))
}

func countRegions(used map[pt]bool) int {
	var count int
	regions := map[pt]*region{}
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			p := pt{x, y}
			if !used[p] {
				continue
			}
			left := pt{x - 1, y}
			up := pt{x, y - 1}
			if regions[left] != nil && regions[up] != nil {
				if regions[left] != regions[up] {
					merge(regions, left, up)
					count--
				}
				regions[p] = regions[left]
				(*regions[p])[p] = true
			} else if regions[left] != nil {
				regions[p] = regions[left]
				(*regions[p])[p] = true
			} else if regions[up] != nil {
				regions[p] = regions[up]
				(*regions[p])[p] = true
			} else {
				regions[p] = &region{p: true}
				count++
			}
		}
	}
	return count
}

func merge(regions map[pt]*region, left, up pt) {
	for upPt := range *regions[up] {
		(*regions[left])[upPt] = true
		regions[upPt] = regions[left]
	}
}

// Slighly modified from day 10 (to give a binary representation)
func denseHash(lengths []byte, rounds int) string {
	sparse := sparseHash(lengths, rounds)
	var hash string
	for block := 0; block < knotSize/16; block++ {
		offset := block * 16
		acc := sparse[offset]
		for i := offset + 1; i < offset+16; i++ {
			acc ^= sparse[i]
		}
		hash += fmt.Sprintf("%08.8s", strconv.FormatInt(int64(acc), 2))
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
