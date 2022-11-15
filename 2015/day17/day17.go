package main

import (
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	containers := util.IntLines("input.txt")

	goal := 150
	count := 0
	depths := map[int]int{}
	for i := 0; i < len(containers); i++ {
		count += countCombinations(containers, goal, i, 0, 0, &depths)
	}
	fmt.Println("part1:", count)

	minDepth := math.MaxInt
	minDepthCount := 0
	for k, v := range depths {
		if k < minDepth {
			minDepth = k
			minDepthCount = v
		}
	}
	fmt.Println("part2:", minDepthCount)
}

func countCombinations(containers []int, remaining int, pos int, count int, depth int, depths *map[int]int) int {
	remaining -= containers[pos]

	if remaining == 0 {
		(*depths)[depth+1]++
		return count + 1
	} else if remaining < 0 {
		return count
	}

	for i := pos + 1; i < len(containers); i++ {
		count = countCombinations(containers, remaining, i, count, depth+1, depths)
	}

	return count
}
