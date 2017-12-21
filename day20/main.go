package main

import (
	"bufio"
	"fmt"
	"os"
)

type vec struct {
	x, y, z int
}

func (a vec) add(b vec) vec {
	return vec{a.x + b.x, a.y + b.y, a.z + b.z}
}

type particle struct {
	p, v, a vec
	id      int
}

func (p *particle) tick() {
	p.v = p.v.add(p.a)
	p.p = p.p.add(p.v)
}

func main() {
	particles := map[*particle]bool{}
	var id int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p := particle{id: id}
		id++
		fmt.Sscanf(
			scanner.Text(),
			"p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&p.p.x, &p.p.y, &p.p.z,
			&p.v.x, &p.v.y, &p.v.z,
			&p.a.x, &p.a.y, &p.a.z,
		)
		particles[&p] = true
	}
	fmt.Printf("part 1: %v\n", closestInLongTerm(particles).id)
	// Simulating for 100 steps gets us the right answer... is there some nice
	// way to know when we're done?
	simulateCollisions(particles, 100)
	fmt.Printf("part 2: %v\n", len(particles))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func manhattan(v vec) int {
	return abs(v.x) + abs(v.y) + abs(v.z)
}

func closestInLongTerm(particles map[*particle]bool) *particle {
	var aMin, vMin, pMin int
	var result *particle
	for particle := range particles {
		a := manhattan(particle.a)
		v := manhattan(particle.v)
		p := manhattan(particle.v)
		if result == nil ||
			a < aMin ||
			a == aMin && v < vMin ||
			a == aMin && v == vMin && p < pMin {
			aMin = a
			vMin = v
			pMin = p
			result = particle
		}
	}
	return result
}

func simulateCollisions(particles map[*particle]bool, steps int) {
	for i := 0; i < steps; i++ {
		space := map[vec]map[*particle]bool{}
		for p := range particles {
			p.tick()
			if _, ok := space[p.p]; !ok {
				space[p.p] = map[*particle]bool{}
			}
			space[p.p][p] = true
		}
		for _, collisions := range space {
			if len(collisions) > 1 {
				for p := range collisions {
					delete(particles, p)
				}
			}
		}
	}
}
