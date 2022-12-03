package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	part1 := 0
	for i, line := range lines {
		r := intersect(line[:len(line)/2], line[len(line)/2:])
		if len(r) != 1 {
			panic(fmt.Sprintf("invalid rucksack %d: %s", i, line))
		}
		part1 += priority(util.FirstRune(r))
	}
	fmt.Println("part1:", part1)

	part2 := 0
	for i := 0; i < len(lines); i += 3 {
		r := intersect(intersect(lines[i], lines[i+1]), lines[i+2])
		if len(r) != 1 {
			panic(fmt.Sprintf("invalid group %d-%d", i, i+2))
		}
		part2 += priority(util.FirstRune(r))
	}
	fmt.Println("part2:", part2)
}

func priority(r rune) int {
	if 'a' <= r && r <= 'z' {
		return int(r-'a') + 1
	} else if 'A' <= r && r <= 'Z' {
		return int(r-'A') + 27
	}
	panic("invalid item: " + string(r))
}

func intersect(a, b string) string {
	var result string
	for _, r := range a {
		if strings.ContainsRune(b, r) && !strings.ContainsRune(result, r) {
			result += string(r)
		}
	}
	return result
}
