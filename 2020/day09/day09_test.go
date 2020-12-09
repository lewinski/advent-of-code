package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Lines("input-example.txt")
	values := ints(input)
	preamble := 5

	got := findInvalidValue(values, preamble)
	if got != 127 {
		t.Errorf("wanted 127, got %d", got)
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Lines("input-example.txt")
	values := ints(input)
	got := findEncryptionWeakness(values, 127)
	if got != 62 {
		t.Errorf("wanted 62, got %d", got)
	}
}
