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

func runProgram(program program) (accum int, booted bool) {
	m := machine{}
	visited := map[int]bool{}

	for {
		_, ok := visited[m.pc]
		if ok {
			return m.accum, false
		}
		visited[m.pc] = true

		if m.pc >= len(program) {
			break
		}

		i := program[m.pc]
		switch i.opcode {
		case "nop":
			m.pc++
			break
		case "acc":
			m.accum += i.argument
			m.pc++
			break
		case "jmp":
			m.pc += i.argument
			break
		}
	}

	return m.accum, true
}

func fixProgram(program program) (accum int, booted bool) {
	for i := len(program) - 1; i > 0; i-- {
		tmp := append([]instruction{}, program...)
		switch tmp[i].opcode {
		case "acc":
			continue
		case "nop":
			tmp[i].opcode = "jmp"
		case "jmp":
		tmp[i].opcode = "nop"
		}
		part2, booted := runProgram(tmp)
		if booted {
			return part2, booted
		}
	}
	panic("couldn't fix it")
}
