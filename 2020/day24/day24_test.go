package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Lines("input-example.txt")
	grid := initGrid(input)

	got := grid.flipped()
	if got != 10 {
		t.Errorf("wanted 10 black tiles, got %d", got)
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Lines("input-example.txt")

	tests := []struct {
		day, want int
	}{
		{1, 15},
		{2, 12},
		{3, 25},
		{4, 14},
		{5, 23},
		{6, 28},
		{7, 41},
		{8, 37},
		{9, 49},
		{10, 37},
		{20, 132},
		{30, 259},
		{40, 406},
		{50, 566},
		{60, 788},
		{70, 1106},
		{80, 1373},
		{90, 1844},
		{100, 2208},
	}

	for _, tt := range tests {
		grid := initGrid(input)
		for i := 0; i < tt.day; i++ {
			grid = step(grid)
		}
		got := grid.flipped()
		if got != tt.want {
			t.Errorf("day %d: wanted %d black tiles, got %d", tt.day, tt.want, got)
		}
	}
}
