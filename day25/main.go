package main

import "fmt"

type state string

type operation struct {
	write, move int
	nextState   state
}

type transition map[int]operation

type machine struct {
	state       state
	cursor      int
	tape        map[int]int
	transitions map[state]transition
}

func (m *machine) step() {
	op := m.transitions[m.state][m.tape[m.cursor]]
	m.tape[m.cursor] = op.write
	m.cursor += op.move
	m.state = op.nextState
}

func main() {
	var steps int
	m := machine{tape: map[int]int{}, transitions: map[state]transition{}}
	fmt.Scanf("Begin in state %1s.\n", &m.state)
	fmt.Scanf("Perform a diagnostic checksum after %d steps.\n", &steps)
	for {
		s, t, done := parseTransition()
		if done {
			break
		}
		m.transitions[s] = t
	}
	for i := 0; i < steps; i++ {
		m.step()
	}
	fmt.Printf("part 1: %v\n", checksum(m.tape))
}

func checksum(tape map[int]int) int {
	var sum int
	for _, n := range tape {
		sum += n
	}
	return sum
}

func parseTransition() (state, transition, bool) {
	var s state
	if _, err := fmt.Scanf("\nIn state %1s:\n", &s); err != nil {
		return s, nil, true
	}
	t := map[int]operation{}
	for i := 0; i < 2; i++ {
		val, op := parseOperation()
		t[val] = op
	}
	return s, t, false
}

func parseOperation() (int, operation) {
	var val int
	var op operation
	var direction string
	fmt.Scanf(" If the current value is %1d:\n", &val)
	fmt.Scanf(" - Write the value %1d.\n", &op.write)
	fmt.Scanf(" - Move one slot to the %s", &direction)
	switch direction {
	case "left.":
		op.move = -1
	case "right.":
		op.move = +1
	}
	fmt.Scanf(" - Continue with state %1s.\n", &op.nextState)
	return val, op
}
