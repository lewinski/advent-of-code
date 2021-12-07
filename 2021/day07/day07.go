package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	max := math.MinInt
	positions := []int{}
	for _, s := range strings.Split(lines[0], ",") {
		pos := util.MustAtoi(s)
		if pos > max {
			max = pos
		}
		positions = append(positions, pos)
	}

	fmt.Println("part1:", part1(max, positions))
	fmt.Println("part2:", part2(max, positions))
}

func part1(max int, positions []int) int {
	bestCost := math.MaxInt
	for curPos := 0; curPos < max; curPos++ {
		curCost := 0
		for _, p := range positions {
			curCost += util.IAbs(p - curPos)
		}
		if curCost < bestCost {
			bestCost = curCost
		}
	}

	return bestCost
}

func part2(max int, positions []int) int {
	bestCost := math.MaxInt
	for curPos := 0; curPos < max; curPos++ {
		curCost := 0
		for _, p := range positions {
			n := util.IAbs(p - curPos)
			curCost += (n * (n + 1)) / 2
		}
		if curCost < bestCost {
			bestCost = curCost
		}
	}

	return bestCost
}
