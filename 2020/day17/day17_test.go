package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Lines("input-example.txt")
	steps := 6
	cube := parseInput1(input, steps)
	for i := 0; i < steps; i++ {
		cube = cube.step()
	}

	count := 0
	for _, x := range cube.cells {
		count += x
	}

	if count != 112 {
		t.Errorf("expected 112 cells, got %d", count)
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Lines("input-example.txt")
	steps := 6
	cube := parseInput2(input, steps)

	for i := 0; i < steps; i++ {
		cube = cube.step()
	}

	count := 0
	for _, x := range cube.cells {
		count += x
	}

	if count != 848 {
		t.Errorf("expected 848 cells, got %d", count)
	}
}
