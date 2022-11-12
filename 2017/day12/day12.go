package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	adj := map[int][]int{}

	for _, line := range lines {
		f := strings.SplitN(line, " <-> ", 2)
		x := util.MustAtoi(f[0])
		adj[x] = []int{}
		for _, c := range strings.Split(f[1], ", ") {
			adj[x] = append(adj[x], util.MustAtoi(c))
		}
	}

	fmt.Println("part1:", part1(adj))
	fmt.Println("part2:", part2(adj))
}

func part1(adj map[int][]int) int {
	seen := map[int]bool{}

	var visit func(x int)
	visit = func(x int) {
		seen[x] = true

		for _, y := range adj[x] {
			if !seen[y] {
				visit(y)
			}
		}
	}

	visit(0)

	return len(seen)
}

func part2(adj map[int][]int) int {
	groups := map[int]int{}
	groupNums := map[int]bool{}

	var visit func(g, x int)
	visit = func(g, x int) {
		groups[x] = g
		groupNums[g] = true

		for _, y := range adj[x] {
			if groups[y] == 0 {
				visit(g, y)
			}
		}
	}

	for len(groups) != len(adj) {
		for g := range adj {
			if groups[g] > 0 {
				continue
			}
			visit(g, g)
			break
		}
	}

	return len(groupNums)
}
