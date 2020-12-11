package main

import (
	"fmt"
	"sort"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	values := util.IntLines("input.txt")
	sort.Ints(values)
	_, diff1, _, diff3 := validChain(values)
	fmt.Println("part1:", diff1*diff3)
	fmt.Println("part2:", totalArrangements(values))
}

func validChain(values []int) (max, diff1, diff2, diff3 int) {
	for _, j := range values {
		switch j - max {
		case 1:
			diff1++
		case 2:
			diff2++
		case 3:
			diff3++
		default:
			return 0, 0, 0, 0
		}
		max = j
	}
	diff3++
	max += 3
	return
}

func totalArrangements(values []int) int {
	// figure out the maximum joltage
	max, _, _, _ := validChain(values)

	// create an easy way to check if joltage is reachable
	reachable := map[int]bool{0: true, max: true}
	for _, v := range values {
		reachable[v] = true
	}

	// count the number of routes there are to each joltage
	routes := map[int]int{}

	// there's one way to 0 jolts, since we start there
	routes[0] = 1

	// for each joltage our adapter chain can reach
	for _, v := range append(values, max) {
		// if we can get to 1 less than the joltage, we can also get to this joltage
		if _, found := reachable[v-1]; found {
			routes[v] += routes[v-1]
		}
		// .. 2 less
		if _, found := reachable[v-2]; found {
			routes[v] += routes[v-2]
		}
		// .. 3 less
		if _, found := reachable[v-3]; found {
			routes[v] += routes[v-3]
		}
	}

	return routes[max]
}
