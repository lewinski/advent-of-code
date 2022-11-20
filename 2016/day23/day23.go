package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type computer struct {
	reg     map[string]int
	pc      int
	program []string
}

func makeComputer(program []string) computer {
	c := computer{
		reg: map[string]int{
			"a": 0,
			"b": 0,
			"c": 0,
			"d": 0,
		},
		pc:      0,
		program: program,
	}
	return c
}

func main() {
	c := makeComputer(util.Lines("input.txt"))
	c.reg["a"] = 7
	c.run()
	fmt.Println("part1:", c.reg["a"])

	// who cares
	// go run day23.go  330.71s user 5.61s system 104% cpu 5:22.93 total
	c = makeComputer(util.Lines("input.txt"))
	c.reg["a"] = 12
	c.run()
	fmt.Println("part2:", c.reg["a"])
}

func (c *computer) run() {
	for {
		if c.pc < 0 || c.pc >= len(c.program) {
			break
		}

		args := strings.Fields(c.program[c.pc])
		switch args[0] {
		case "cpy":
			var value int
			if v, ok := c.reg[args[1]]; ok {
				value = v
			} else {
				value = util.MustAtoi(args[1])
			}
			if _, found := c.reg[args[2]]; found {
				c.reg[args[2]] = value
			}
		case "inc":
			if _, ok := c.reg[args[1]]; ok {
				c.reg[args[1]]++
			}
		case "dec":
			if _, ok := c.reg[args[1]]; ok {
				c.reg[args[1]]--
			}
		case "jnz":
			var jump int
			if v, ok := c.reg[args[2]]; ok {
				jump = v
			} else {
				jump = util.MustAtoi(args[2])
			}

			if v, ok := c.reg[args[1]]; ok {
				if v != 0 {
					c.pc += jump
					continue
				}
			} else if util.MustAtoi(args[1]) != 0 {
				c.pc += jump
				continue
			}
		case "tgl":
			var target int
			if v, ok := c.reg[args[1]]; ok {
				target = c.pc + v
			} else {
				target = c.pc + util.MustAtoi(args[1])
			}
			if target < 0 || target >= len(c.program) {
				break
			}
			args := strings.Fields(c.program[target])
			switch len(args) {
			case 2:
				if args[0] == "inc" {
					args[0] = "dec"
				} else {
					args[0] = "inc"
				}
			case 3:
				if args[0] == "jnz" {
					args[0] = "cpy"
				} else {
					args[0] = "jnz"
				}
			}
			c.program[target] = strings.Join(args, " ")
		}
		c.pc++
	}
}
