package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	bursts1 = 10000
	bursts2 = 10000000
)

type vec struct {
	x, y int
}

type virusCarrier struct {
	position, direction vec
}

type state int

const (
	clean state = iota
	weakened
	infected
	flagged
)

type grid struct {
	states map[vec]state
}

func newGrid() *grid {
	return &grid{map[vec]state{}}
}

func (vc *virusCarrier) move() {
	vc.position = vec{
		vc.position.x + vc.direction.x,
		vc.position.y + vc.direction.y,
	}
}

func (vc *virusCarrier) turnRight() {
	vc.direction = vec{-vc.direction.y, vc.direction.x}
}

func (vc *virusCarrier) turnLeft() {
	vc.direction = vec{vc.direction.y, -vc.direction.x}
}

func (g *grid) weaken(v vec) {
	g.states[v] = weakened
}

func (g *grid) infect(v vec) {
	g.states[v] = infected
}

func (g *grid) flag(v vec) {
	g.states[v] = flagged
}

func (g *grid) clean(v vec) {
	delete(g.states, v)
}

func (g *grid) state(v vec) state {
	return g.states[v]
}

func main() {
	grid1 := newGrid()
	grid2 := newGrid()
	var center vec
	scanner := bufio.NewScanner(os.Stdin)
	for y := 0; scanner.Scan(); y++ {
		for x, c := range scanner.Text() {
			if c == '#' {
				grid1.infect(vec{x, y})
				grid2.infect(vec{x, y})
			}
			center = vec{x / 2, y / 2}
		}
	}
	vc1 := virusCarrier{center, vec{0, -1}}
	var infectionCount1 int
	for burst := 0; burst < bursts1; burst++ {
		switch grid1.state(vc1.position) {
		case clean:
			vc1.turnLeft()
			grid1.infect(vc1.position)
			infectionCount1++
		case infected:
			vc1.turnRight()
			grid1.clean(vc1.position)
		}
		vc1.move()
	}
	fmt.Printf("part 1: %v\n", infectionCount1)
	vc2 := virusCarrier{center, vec{0, -1}}
	var infectionCount2 int
	for burst := 0; burst < bursts2; burst++ {
		switch grid2.state(vc2.position) {
		case clean:
			vc2.turnLeft()
			grid2.weaken(vc2.position)
		case weakened:
			grid2.infect(vc2.position)
			infectionCount2++
		case infected:
			grid2.flag(vc2.position)
			vc2.turnRight()
		case flagged:
			grid2.clean(vc2.position)
			vc2.turnRight()
			vc2.turnRight()
		}
		vc2.move()
	}
	fmt.Printf("part 2: %v\n", infectionCount2)
}
