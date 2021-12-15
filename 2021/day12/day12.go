package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	adjacency := map[string][]string{}

	for _, line := range lines {
		f := strings.SplitN(line, "-", 2)
		adjacency[f[0]] = append(adjacency[f[0]], f[1])
		adjacency[f[1]] = append(adjacency[f[1]], f[0])
	}

	fmt.Println("part1:", part1(adjacency))
	fmt.Println("part2:", part2(adjacency))
}

func isBig(cave string) bool {
	return strings.ToUpper(cave) == cave
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func move(cur []string, next string) []string {
	new := make([]string, len(cur)+1)
	copy(new, cur)
	new[len(new)-1] = next
	return new
}

func smallTwiceOk(path []string) bool {
	seen := map[string]int{}
	for _, s := range path {
		if s == "start" || isBig(s) {
			continue
		}
		seen[s]++
	}
	twice := 0
	for _, c := range seen {
		if c > 2 {
			return false
		} else if c == 2 {
			twice++
		}
	}
	return twice <= 1
}

func part1(adj map[string][]string) int {
	paths := [][]string{}

	stack := [][]string{}
	stack = append(stack, []string{"start"})

	for len(stack) > 0 {
		var cur []string
		cur, stack = stack[0], stack[1:]

		for _, next := range adj[cur[len(cur)-1]] {
			newPath := move(cur, next)
			if next == "start" {
				continue
			} else if next == "end" {
				paths = append(paths, newPath)
			} else if isBig(next) {
				stack = append(stack, newPath)
			} else if !contains(cur, next) {
				stack = append(stack, newPath)
			}
		}

	}

	return len(paths)
}

func part2(adj map[string][]string) int {
	paths := [][]string{}

	stack := [][]string{}
	stack = append(stack, []string{"start"})

	for len(stack) > 0 {
		var cur []string
		cur, stack = stack[0], stack[1:]

		for _, next := range adj[cur[len(cur)-1]] {
			newPath := move(cur, next)
			if next == "start" {
				continue
			} else if next == "end" {
				paths = append(paths, newPath)
			} else if isBig(next) {
				stack = append(stack, newPath)
			} else if !contains(cur, next) {
				stack = append(stack, newPath)
			} else if smallTwiceOk(newPath) {
				stack = append(stack, newPath)
			}
		}
	}

	return len(paths)
}
