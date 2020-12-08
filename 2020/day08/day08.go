package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	program := parseProgram(util.Lines("input.txt"))

	part1, _ := runProgram(program)
	fmt.Println("part1:", part1)

	part2, _ := fixProgram(program)
	fmt.Println("part2:", part2)
}

type instruction struct {
	opcode   string
	argument int
}
type program []instruction

func parseProgram(lines []string) program {
	program := make([]instruction, len(lines))
	for i, line := range lines {
		pair := strings.SplitN(line, " ", 2)
		if len(pair) != 2 {
			continue
		}
		program[i].opcode = pair[0]
		program[i].argument, _ = strconv.Atoi(pair[1])
	}
	return program
}

type machine struct {
	pc    int
	accum int
}

func runProgram(program program) (int, bool) {
	m := machine{}
	visited := map[int]bool{}
	looped := false

	for {
		_, found := visited[m.pc]
		if found {
			looped = true
			return m.accum, looped
		}
		visited[m.pc] = true

		if m.pc >= len(program) {
			return m.accum, looped
		}

		i := program[m.pc]
		switch i.opcode {
		case "nop":
			m.pc++
		case "acc":
			m.accum += i.argument
			m.pc++
		case "jmp":
			m.pc += i.argument
		}
	}
}

func fixProgram(program program) (int, bool) {
	for i := len(program) - 1; i > 0; i-- {
		opcode := program[i].opcode
		switch opcode {
		case "acc":
			continue
		case "nop":
			program[i].opcode = "jmp"
		case "jmp":
			program[i].opcode = "nop"
		}
		accum, looped := runProgram(program)
		program[i].opcode = opcode
		if !looped {
			return accum, looped
		}
	}
	panic("couldn't fix it")
}
