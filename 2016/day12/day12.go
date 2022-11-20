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
	c.run()
	fmt.Println("part1:", c.reg["a"])

	c = makeComputer(util.Lines("input.txt"))
	c.reg["c"] = 1
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
		}
		c.pc++
	}
}
