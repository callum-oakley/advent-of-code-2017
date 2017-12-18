package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type prog struct {
	regs               map[string]int
	i, sendCount, echo int
	blocked            bool
	partner            *prog
	done               chan struct{}
	in, recover        chan int
}

func (p *prog) resolve(identifier string) int {
	if n, err := strconv.Atoi(identifier); err == nil {
		return n
	}
	return p.regs[identifier]
}

func (p *prog) run(instructions []func(*prog)) {
	for p.i >= 0 && p.i < len(instructions) && !p.blocked {
		instructions[p.i](p)
	}
	p.done <- struct{}{}
}

func main() {
	var instructions []func(*prog)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, parse(line))
	}
	p := &prog{
		regs:    map[string]int{},
		recover: make(chan int),
		done:    make(chan struct{}),
	}
	p0 := &prog{
		regs: map[string]int{"p": 0},
		in:   make(chan int, 99),
		done: make(chan struct{}),
	}
	p1 := &prog{
		regs: map[string]int{"p": 1},
		in:   make(chan int, 99),
		done: make(chan struct{}),
	}
	p0.partner, p1.partner = p1, p0
	go p.run(instructions)
	go p0.run(instructions)
	go p1.run(instructions)
	fmt.Printf("part 1: %v\n", <-p.recover)
	<-p0.done
	<-p1.done
	fmt.Printf("part 2: %v\n", p1.sendCount)
}

func parse(line string) func(*prog) {
	fields := strings.Fields(line)
	switch fields[0] {
	case "snd":
		return func(p *prog) {
			if p.partner != nil {
				p.partner.in <- p.resolve(fields[1])
			} else {
				p.echo = p.resolve(fields[1])
			}
			p.sendCount++
			p.i++
		}
	case "set":
		return func(p *prog) {
			p.regs[fields[1]] = p.resolve(fields[2])
			p.i++
		}
	case "add":
		return func(p *prog) {
			p.regs[fields[1]] += p.resolve(fields[2])
			p.i++
		}
	case "mul":
		return func(p *prog) {
			p.regs[fields[1]] *= p.resolve(fields[2])
			p.i++
		}
	case "mod":
		return func(p *prog) {
			p.regs[fields[1]] %= p.resolve(fields[2])
			p.i++
		}
	case "rcv":
		return func(p *prog) {
			if p.partner != nil {
				for {
					select {
					case p.regs[fields[1]] = <-p.in:
						p.blocked = false
						p.i++
						return
					default:
						p.blocked = true
						if p.partner.blocked == true {
							return
						}
					}
				}
			} else if p.regs[fields[1]] != 0 {
				p.recover <- p.echo
				p.blocked = true
			}
		}
	case "jgz":
		return func(p *prog) {
			if p.resolve(fields[1]) > 0 {
				p.i += p.resolve(fields[2])
			} else {
				p.i++
			}
		}
	}
	return nil
}
