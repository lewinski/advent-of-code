package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")[0]
	fmt.Println("part1:", countSafe(input, 40))
	fmt.Println("part1:", countSafe(input, 400000))
}

func countSafe(input string, rows int) int {
	safe := 0
	for i := 0; i < rows; i++ {
		safe += strings.Count(input, ".")
		input = nextTraps(input)
	}
	return safe
}

func nextTraps(last string) string {
	next := make([]byte, len(last))
	last = "." + last + "."
	for i := 1; i < len(last)-1; i++ {
		switch last[i-1 : i+2] {
		case "^^.", ".^^", "^..", "..^":
			next[i-1] = '^'
		default:
			next[i-1] = '.'
		}
	}
	return string(next)
}
