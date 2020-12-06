package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	records := util.Records("input-example.txt")
	groups := parseQuestions(records)
	want := 11
	got := part1(groups)
	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestExamplePart2(t *testing.T) {
	records := util.Records("input-example.txt")
	groups := parseQuestions(records)
	want := 6
	got := part2(groups)
	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}
