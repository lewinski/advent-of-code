package main

import (
	"fmt"
	"sort"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func traverse(line string) (score int, stack []rune) {
	for _, r := range line {
		switch r {
		case '(':
			stack = append(stack, r)
		case '[':
			stack = append(stack, r)
		case '{':
			stack = append(stack, r)
		case '<':
			stack = append(stack, r)
		case ')':
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				score = 3
				return
			}
		case ']':
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				score = 57
				return
			}
		case '}':
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				score = 1197
				return
			}
		case '>':
			if stack[len(stack)-1] == '<' {
				stack = stack[:len(stack)-1]
			} else {
				score = 25137
				return
			}

		}
	}
	return
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		score, _ := traverse(line)
		sum += score
	}
	return sum
}

func autocomplete(remainder []rune) int {
	points := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	x := 0
	for i := len(remainder) - 1; i >= 0; i-- {
		x = (x * 5) + points[remainder[i]]
	}
	return x
}

func part2(lines []string) int {
	autocompletes := []int{}

	for _, line := range lines {
		score, remainder := traverse(line)
		if score == 0 {
			// incomplete
			autocompletes = append(autocompletes, autocomplete(remainder))
		}
	}

	sort.Ints(autocompletes)

	return autocompletes[len(autocompletes)/2]
}
