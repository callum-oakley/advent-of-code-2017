package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type pixel bool
type block [][]pixel

const startingBlock = ".#./..#/###"

func (b block) subBlock(bx, by, size int) block {
	result := make(block, size)
	for y := 0; y < size; y++ {
		result[y] = make([]pixel, size)
		for x := 0; x < size; x++ {
			result[y][x] = b[size*by+y][size*bx+x]
		}
	}
	return result
}

func (b block) hash() string {
	return fmt.Sprintf("%v", b)
}

func main() {
	rulebook := map[string]block{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		from := parseBlock(fields[0])
		to := parseBlock(fields[2])
		for _, tfrom := range transform(from) {
			rulebook[tfrom.hash()] = to
		}
	}
	b := parseBlock(startingBlock)
	var b5 block
	for i := 1; i <= 18; i++ {
		b = enhance(rulebook, b)
		if i == 5 {
			b5 = b
		}
	}
	fmt.Printf("part 1: %v\n", countOn(b5))
	fmt.Printf("part 2: %v\n", countOn(b))
}

func enhance(rulebook map[string]block, b block) block {
	blocks := chunk(b)
	for y := 0; y < len(blocks); y++ {
		for x := 0; x < len(blocks); x++ {
			blocks[y][x] = rulebook[blocks[y][x].hash()]
		}
	}
	return assemble(blocks)
}

func chunk(b block) [][]block {
	var blocks [][]block
	var size int
	if len(b)%2 == 0 {
		size = 2
	} else if len(b)%3 == 0 {
		size = 3
	}
	blocks = make([][]block, len(b)/size)
	for y := 0; y < len(b)/size; y++ {
		blocks[y] = make([]block, len(b)/size)
		for x := 0; x < len(b)/size; x++ {
			blocks[y][x] = b.subBlock(x, y, size)
		}
	}
	return blocks
}

func assemble(blocks [][]block) block {
	size := len(blocks[0][0])
	b := make(block, size*len(blocks))
	for y := 0; y < size*len(blocks); y++ {
		b[y] = make([]pixel, size*len(blocks))
		for x := 0; x < size*len(blocks); x++ {
			b[y][x] = blocks[y/size][x/size][y%size][x%size]
		}
	}
	return b
}

func parseBlock(s string) block {
	rows := strings.Split(s, "/")
	b := make(block, len(rows))
	for y := 0; y < len(rows); y++ {
		b[y] = make([]pixel, len(rows))
		for x := 0; x < len(rows); x++ {
			b[y][x] = rows[y][x] == '#'
		}
	}
	return b
}

func transform(b block) []block {
	var blocks []block
	for _, rb := range rotations(b) {
		blocks = append(blocks, rb, flip(rb))
	}
	return blocks
}

func rotations(b0 block) []block {
	b1 := rotate(b0)
	b2 := rotate(b1)
	b3 := rotate(b2)
	return []block{b0, b1, b2, b3}
}

func rotate(b block) block {
	result := make(block, len(b))
	for y := 0; y < len(b); y++ {
		result[y] = make([]pixel, len(b))
		for x := 0; x < len(b); x++ {
			result[y][x] = b[x][len(b)-1-y]
		}
	}
	return result
}

func flip(b block) block {
	result := make(block, len(b))
	for y := 0; y < len(b); y++ {
		result[y] = make([]pixel, len(b))
		for x := 0; x < len(b); x++ {
			result[y][x] = b[x][y]
		}
	}
	return result
}

func countOn(b block) int {
	var count int
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b); x++ {
			if b[y][x] {
				count++
			}
		}
	}
	return count
}

// Just for fun (unfortunately it doesn't end up very pretty...)
func draw(title string, b block) error {
	img := image.NewRGBA(image.Rect(0, 0, len(b), len(b)))
	for y := 0; y < len(b); y++ {
		for x := 0; x < len(b); x++ {
			if b[y][x] {
				img.Set(x, y, color.RGBA{255, 255, 255, 255})
			} else {
				img.Set(x, y, color.RGBA{})
			}
		}
	}
	f, err := os.OpenFile(
		fmt.Sprintf("%v.png", title),
		os.O_WRONLY|os.O_CREATE,
		0600,
	)
	if err != nil {
		return err
	}
	defer f.Close()
	png.Encode(f, img)
	return nil
}
