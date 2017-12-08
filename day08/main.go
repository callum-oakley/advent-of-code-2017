package main

import (
	"fmt"
	"io"
	"log"
)

type registers map[string]int

type instruction struct {
	reg  string
	val  int
	cond func(registers) bool
}

func main() {
	instructions, err := parse()
	if err != nil {
		log.Fatal(err)
	}
	regs := map[string]int{}
	rollingMax := 0
	for _, ins := range instructions {
		if ins.cond(regs) {
			regs[ins.reg] += ins.val
			if regs[ins.reg] > rollingMax {
				rollingMax = regs[ins.reg]
			}
		}
	}
	fmt.Printf("part 1: %v\n", maxVal(regs))
	fmt.Printf("part 2: %v\n", rollingMax)
}

func maxVal(m map[string]int) int {
	var max int
	initialized := false
	for _, v := range m {
		if !initialized {
			max = v
			initialized = true
			continue
		}
		if v > max {
			max = v
		}
	}
	return max
}

func parse() ([]instruction, error) {
	var instructions []instruction
	for {
		var ins instruction
		var op, condReg, condOp string
		var condVal int
		n, err := fmt.Scanf(
			"%v %v %v if %v %v %v",
			&ins.reg,
			&op,
			&ins.val,
			&condReg,
			&condOp,
			&condVal,
		)
		if n != 6 || err != nil {
			if err != io.EOF {
				return nil, err
			}
			return instructions, nil
		}
		if op == "dec" {
			ins.val *= -1
		}
		switch condOp {
		case "<":
			ins.cond = func(reg registers) bool {
				return reg[condReg] < condVal
			}
		case ">":
			ins.cond = func(reg registers) bool {
				return reg[condReg] > condVal
			}
		case "<=":
			ins.cond = func(reg registers) bool {
				return reg[condReg] <= condVal
			}
		case ">=":
			ins.cond = func(reg registers) bool {
				return reg[condReg] >= condVal
			}
		case "==":
			ins.cond = func(reg registers) bool {
				return reg[condReg] == condVal
			}
		case "!=":
			ins.cond = func(reg registers) bool {
				return reg[condReg] != condVal
			}
		default:
			return nil, fmt.Errorf("Unsusported op: '%v'", condOp)
		}
		instructions = append(instructions, ins)
	}
}
