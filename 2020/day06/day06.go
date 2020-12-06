package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	groups := parseQuestions(util.Records("input.txt"))

	fmt.Println("part1:", part1(groups))
	fmt.Println("part2:", part2(groups))
}

type group struct {
	people    int
	questions map[rune]int
}

func newGroup() group {
	g := group{}
	g.questions = make(map[rune]int)
	return g
}

func parseQuestions(records []string) []group {
	groups := make([]group, 0, len(records))

	for _, record := range records {
		g := newGroup()
		for _, line := range strings.Split(record, "\n") {
			for _, c := range line {
				g.questions[c]++
			}
			g.people++
		}
		groups = append(groups, g)
	}
	return groups
}

func part1(groups []group) (total int) {
	for _, g := range groups {
		total += len(g.questions)
	}
	return
}

func part2(groups []group) (total int) {
	for _, g := range groups {
		for _, c := range g.questions {
			if c == g.people {
				total++
			}
		}
	}
	return
}
