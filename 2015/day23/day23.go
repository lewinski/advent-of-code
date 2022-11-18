package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", runProgram(lines, 0))
	fmt.Println("part2:", runProgram(lines, 1))
}

func runProgram(lines []string, start int) int {
	var a, b, pc int

	a = start

	for {
		if pc >= len(lines) {
			return b
		}

		switch lines[pc] {
		case "hlf a":
			a /= 2
			pc++

		case "inc a":
			a++
			pc++
		case "inc b":
			b++
			pc++

		case "jie a, +4":
			if a%2 == 0 {
				pc += 4
			} else {
				pc++
			}

		case "jio a, +16":
			if a == 1 {
				pc += 16
			} else {
				pc++
			}
		case "jio a, +8":
			if a == 1 {
				pc += 8
			} else {
				pc++
			}

		case "jmp +2":
			pc += 2
		case "jmp +23":
			pc += 23
		case "jmp -7":
			pc -= 7

		case "tpl a":
			a *= 3
			pc++

		default:
			panic("unknown instruction: " + lines[pc])
		}
	}
}
