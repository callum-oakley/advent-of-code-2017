package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type prog struct {
	regs        map[string]int
	i, mulCount int
}

func (p *prog) resolve(identifier string) int {
	if n, err := strconv.Atoi(identifier); err == nil {
		return n
	}
	return p.regs[identifier]
}

func main() {
	var instructions []func(*prog)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, parse(line))
	}
	p1 := &prog{
		regs: map[string]int{},
	}
	for p1.i >= 0 && p1.i < len(instructions) {
		instructions[p1.i](p1)
	}
	fmt.Printf("part 1: %v\n", p1.mulCount)
	p2 := &prog{
		regs: map[string]int{"a": 1},
	}
	for p2.i >= 0 && p2.i < len(instructions) {
		// Opimisation based on the observation that the block of code from
		// instruction 8 to 25 is determining if register b holds a prime, and
		// incrementing h if not.
		if p2.i == 8 {
			if !isPrime(p2.regs["b"]) {
				p2.regs["h"]++
			}
			p2.i = 26
		} else {
			instructions[p2.i](p2)
		}
	}
	fmt.Printf("part 2: %v\n", p2.regs["h"])
}

func isPrime(n int) bool {
	for m := 2; m <= int(math.Sqrt(float64(n))); m++ {
		if n%m == 0 {
			return false
		}
	}
	return true
}

func parse(line string) func(*prog) {
	fields := strings.Fields(line)
	switch fields[0] {
	case "set":
		return func(p *prog) {
			p.regs[fields[1]] = p.resolve(fields[2])
			p.i++
		}
	case "sub":
		return func(p *prog) {
			p.regs[fields[1]] -= p.resolve(fields[2])
			p.i++
		}
	case "mul":
		return func(p *prog) {
			p.regs[fields[1]] *= p.resolve(fields[2])
			p.i++
			p.mulCount++
		}
	case "jnz":
		return func(p *prog) {
			if p.resolve(fields[1]) != 0 {
				p.i += p.resolve(fields[2])
			} else {
				p.i++
			}
		}
	}
	return nil
}
