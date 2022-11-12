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

func part1(input string) int {
	sum := 0
	last := util.LastRune(input)
	for _, r := range input {
		if r == last {
			sum += int(r - '0')
		}
		last = r

	}
	return sum
}

func part2(input string) int {
	sum := 0
	half := len(input) / 2
	for i := 0; i < half; i++ {
		if input[i] == input[i+half] {
			sum += 2 * int(input[i]-'0')
		}
	}
	return sum
}
