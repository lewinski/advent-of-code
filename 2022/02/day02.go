package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func part1(lines []string) int {
	scores := map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	total := 0
	for _, line := range lines {
		total += scores[line]
	}
	return total
}

func part2(lines []string) int {
	scores := map[string]int{
		// losses
		"A X": 3 + 0,
		"B X": 1 + 0,
		"C X": 2 + 0,

		// ties
		"A Y": 1 + 3,
		"B Y": 2 + 3,
		"C Y": 3 + 3,

		// wins
		"A Z": 2 + 6,
		"B Z": 3 + 6,
		"C Z": 1 + 6,
	}

	total := 0
	for _, line := range lines {
		total += scores[line]
	}
	return total
}
