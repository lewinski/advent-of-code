package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	program := util.IntLines("input.txt")
	fmt.Println("part1:", part1(program))
	program = util.IntLines("input.txt")
	fmt.Println("part2:", part2(program))
}

func part1(program []int) int {
	steps := 0

	pos := 0
	for {
		next := pos + program[pos]
		program[pos]++
		steps++
		if next < 0 || next >= len(program) {
			break
		}
		pos = next
	}

	return steps
}

func part2(program []int) int {
	steps := 0

	pos := 0
	for {
		next := pos + program[pos]
		if program[pos] >= 3 {
			program[pos]--
		} else {
			program[pos]++
		}
		steps++
		if next < 0 || next >= len(program) {
			break
		}
		pos = next
	}

	return steps
}
