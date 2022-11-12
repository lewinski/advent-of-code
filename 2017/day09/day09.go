package main

import (
	"fmt"
	"regexp"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", part1(lines[0]))
	fmt.Println("part2:", part2(lines[0]))
}

func part1(groups string) int {
	groups = knockoutIgnores(groups)
	groups = regexp.MustCompile("<[^>]*?>").ReplaceAllString(groups, "<g>")

	accum := 0
	depth := 0
	for _, r := range groups {
		if r == '{' {
			depth++
			accum += depth
		} else if r == '}' {
			depth--
		}
	}
	return accum
}

func part2(groups string) int {
	groups = knockoutIgnores(groups)

	accum := 0
	regexp.MustCompile("<[^>]*?>").ReplaceAllStringFunc(groups, func(garbage string) string {
		accum += len(garbage) - 2
		return "<>"
	})
	return accum
}

func knockoutIgnores(s string) string {
	return regexp.MustCompile("!.").ReplaceAllString(s, "")
}
