package main

import (
	"strings"
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Records("input-example.txt")

	rules := parseRules(input[0])
	lines := strings.Split(input[1], "\n")

	matches := countPart1Matches(lines, rules)
	if matches != 2 {
		t.Errorf("wanted 2 matches, got %d", matches)
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Records("input-example2.txt")

	rules := parseRules(input[0])
	lines := strings.Split(input[1], "\n")

	part1 := countPart1Matches(lines, rules)
	part2 := countPart2Matches(lines, rules)
	if part1 != 3 {
		t.Errorf("wanted 3 part1 matches, got %d", part1)
	}
	if part2 != 12 {
		t.Errorf("wanted 12 part2 matches, got %d", part2)
	}
}
