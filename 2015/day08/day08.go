package main

import (
	"fmt"
	"regexp"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input2.txt")

	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))

}

func part1(lines []string) int {
	code := 0
	memory := 0

	escapes := regexp.MustCompile(`\\\\|\\"|\\x[0-9a-f]{2}`)

	for _, l := range lines {
		matches := escapes.FindAllString(l, -1)

		code += len(l)
		memory += len(l) - 2
		for _, m := range matches {
			memory -= len(m) - 1
		}
	}

	return code - memory
}

func part2(lines []string) int {
	code := 0
	memory := 0

	escapes := regexp.MustCompile(`\\|"`)

	for _, l := range lines {
		matches := escapes.FindAllString(l, -1)

		memory += len(l)
		code += len(l) + 2 + len(matches)
	}

	return code - memory
}
