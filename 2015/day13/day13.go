package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	start := ""
	deltas := map[string]int{}
	names := map[string]bool{}

	for _, line := range lines {
		words := strings.Split(strings.TrimRight(line, "."), " ")
		delta := util.MustAtoi(words[3])
		if words[2] == "lose" {
			delta = -delta
		}
		n1 := words[0]
		n2 := words[10]

		names[n1] = true
		names[n2] = true
		if start == "" {
			start = n1
		}

		// store both directions to be lazy
		deltas[n1+n2] += delta
		deltas[n2+n1] += delta
	}

	// since it is a circle, we can just always start with someone
	fmt.Println("part1: ", visit(start, start, 0, map[string]bool{start: true}, names, deltas))

	names["me"] = true
	fmt.Println("part2: ", visit(start, start, 0, map[string]bool{start: true}, names, deltas))
}

func visit(first string, last string, delta int, seen map[string]bool, names map[string]bool, deltas map[string]int) int {
	if len(seen) == len(names) {
		return delta + deltas[first+last]
	}

	max := math.MinInt
	for name := range names {
		if seen[name] {
			continue
		}

		d := deltas[name+last]

		seen[name] = true
		max = util.IMax(max, visit(first, name, delta+d, seen, names, deltas))
		delete(seen, name)
	}
	return max
}
