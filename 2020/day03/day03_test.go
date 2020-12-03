package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExampleTrees(t *testing.T) {
	tests := []struct {
		right, down, want int
	}{
		{3, 1, 7},
		{1, 1, 2},
		{5, 1, 3},
		{7, 1, 4},
		{1, 2, 2},
	}

	forest := util.Lines("input-example.txt")

	for i, tt := range tests {
		got := trees(forest, tt.right, tt.down)
		if got != tt.want {
			t.Errorf("test %d: slope = %d right, %d down; want %d trees, got %d", i, tt.right, tt.down, tt.want, got)
		}
	}
}
