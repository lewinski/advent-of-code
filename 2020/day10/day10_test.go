package main

import (
	"sort"
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExample1Part1(t *testing.T) {
	values := util.IntLines("input-example.txt")
	sort.Ints(values)
	joltage, diff1, diff2, diff3 := validChain(values)
	if joltage != 22 {
		t.Errorf("wanted max joltage of 22, got %d", joltage)
	}
	if diff1 != 7 {
		t.Errorf("wanted 1-differences of 7, got %d", diff1)
	}
	if diff2 != 0 {
		t.Errorf("wanted 2-differences of 0, got %d", diff2)
	}
	if diff3 != 5 {
		t.Errorf("wanted 3-differences of 5, got %d", diff3)
	}
}

func TestExample2Part1(t *testing.T) {
	values := util.IntLines("input-example2.txt")
	sort.Ints(values)
	joltage, diff1, diff2, diff3 := validChain(values)
	if joltage != 52 {
		t.Errorf("wanted max joltage of 52, got %d", joltage)
	}
	if diff1 != 22 {
		t.Errorf("wanted 1-differences of 22, got %d", diff1)
	}
	if diff2 != 0 {
		t.Errorf("wanted 2-differences of 0, got %d", diff2)
	}
	if diff3 != 10 {
		t.Errorf("wanted 3-differences of 10, got %d", diff3)
	}
}

func TestExample1Part2(t *testing.T) {
	values := util.IntLines("input-example.txt")
	sort.Ints(values)
	arrangements := totalArrangements(values)
	if arrangements != 8 {
		t.Errorf("wanted 8 arrangements, got %d", arrangements)
	}
}

func TestExample2Part2(t *testing.T) {
	values := util.IntLines("input-example2.txt")
	sort.Ints(values)
	arrangements := totalArrangements(values)
	if arrangements != 19208 {
		t.Errorf("wanted 19208 arrangements, got %d", arrangements)
	}
}
