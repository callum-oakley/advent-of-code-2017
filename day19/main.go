package main

import (
	"bufio"
	"fmt"
	"os"
)

type vec struct {
	x, y int
}

func (v vec) turnLeft() vec {
	return vec{-v.y, v.x}
}

func (v vec) turnRight() vec {
	return vec{v.y, -v.x}
}

func (a vec) add(b vec) vec {
	return vec{a.x + b.x, a.y + b.y}
}

type packet struct {
	position, direction vec
	collection          string
}

func (p *packet) canTravel(maze []string, direction vec) bool {
	nextPos := p.position.add(direction)
	r := maze[nextPos.y][nextPos.x]
	if direction.x == 0 && r == '|' {
		return true
	}
	if direction.y == 0 && r == '-' {
		return true
	}
	return false
}

func (p *packet) navigate(maze []string) int {
	for steps := 0; ; steps++ {
		switch r := maze[p.position.y][p.position.x]; r {
		case '|', '-':
		case '+':
			left := p.direction.turnLeft()
			right := p.direction.turnRight()
			if p.canTravel(maze, left) {
				p.direction = left
			} else if p.canTravel(maze, right) {
				p.direction = right
			} else {
				return steps
			}
		case ' ':
			return steps
		default:
			p.collection += string(r)
		}
		p.position = p.position.add(p.direction)
	}
}

func main() {
	var maze []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		maze = append(maze, scanner.Text())
	}
	p := packet{position: findStart(maze), direction: vec{0, 1}}
	steps := p.navigate(maze)
	fmt.Printf("part 1: %v\n", p.collection)
	fmt.Printf("part 2: %v\n", steps)
}

func findStart(maze []string) vec {
	for x, r := range maze[0] {
		if r == '|' {
			return vec{x, 0}
		}
	}
	return vec{}
}
