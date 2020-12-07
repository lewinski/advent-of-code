package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Lines("input-example.txt")
	bags := parseBags(input)
	got := bagsContaining("shiny gold", bags)

	if len(got) != 4 {
		t.Errorf("wanted 4 bags, got %d", len(got))
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Lines("input-example.txt")
	bags := parseBags(input)

	got := 0
	for _, c := range contentsOf("shiny gold", bags) {
		got += c
	}

	if got != 32 {
		t.Errorf("wanted 32 bags, got %d", got)
	}
}
