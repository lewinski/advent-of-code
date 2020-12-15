package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	tests := []struct {
		seed  []int
		turns int
		want  int
	}{
		{[]int{0, 3, 6}, 1, 0},
		{[]int{0, 3, 6}, 2, 3},
		{[]int{0, 3, 6}, 3, 6},
		{[]int{0, 3, 6}, 4, 0},
		{[]int{0, 3, 6}, 5, 3},
		{[]int{0, 3, 6}, 6, 3},
		{[]int{0, 3, 6}, 7, 1},
		{[]int{0, 3, 6}, 8, 0},
		{[]int{0, 3, 6}, 9, 4},
		{[]int{0, 3, 6}, 10, 0},
		{[]int{0, 3, 6}, 2020, 436},
		{[]int{1, 3, 2}, 2020, 1},
		{[]int{2, 1, 3}, 2020, 10},
		{[]int{1, 2, 3}, 2020, 27},
		{[]int{2, 3, 1}, 2020, 78},
		{[]int{3, 2, 1}, 2020, 438},
		{[]int{3, 1, 2}, 2020, 1836},
	}

	for i, tt := range tests {
		got := elfGame(tt.seed, tt.turns)
		if got != tt.want {
			t.Errorf("%d: wanted %d, got %d", i, tt.want, got)
		}
	}
}

func TestExamplePart2(t *testing.T) {
	tests := []struct {
		seed  []int
		turns int
		want  int
	}{
		{[]int{0, 3, 6}, 30000000, 175594},
		{[]int{1, 3, 2}, 30000000, 2578},
		{[]int{2, 1, 3}, 30000000, 3544142},
		{[]int{1, 2, 3}, 30000000, 261214},
		{[]int{2, 3, 1}, 30000000, 6895259},
		{[]int{3, 2, 1}, 30000000, 18},
		{[]int{3, 1, 2}, 30000000, 362},
	}

	for i, tt := range tests {
		got := elfGame(tt.seed, tt.turns)
		if got != tt.want {
			t.Errorf("%d: wanted %d, got %d", i, tt.want, got)
		}
	}
}
