package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.IntLines("input.txt")
	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func part1(depths []int) int {
	incs := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			incs++
		}
	}
	return incs
}

func part2(depths []int) int {
	windows := make([]int, len(depths)-2)
	for i := 0; i < len(depths)-2; i++ {
		windows[i] = depths[i] + depths[i+1] + depths[i+2]
	}
	return part1(windows)
}
