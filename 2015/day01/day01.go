package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")
	fmt.Println("part1:", part1(lines[0]))
	fmt.Println("part2:", part2(lines[0]))
}

func part1(directions string) int {
	floor := 0
	for _, c := range directions {
		if c == '(' {
			floor++;
		} else if c == ')' {
			floor--;
		}
	}
	return floor
}

func part2(directions string) int {
	floor := 0
	for i, c := range directions {
		if c == '(' {
			floor++;
		} else if c == ')' {
			floor--;
		}
		if floor < 0 {
			return i + 1
		}
	}
	panic("oh no")
}
