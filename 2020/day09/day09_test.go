package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	values := util.IntLines("input-example.txt")
	preamble := 5

	got := findInvalidValue(values, preamble)
	if got != 127 {
		t.Errorf("wanted 127, got %d", got)
	}
}

func TestExamplePart2(t *testing.T) {
	values := util.IntLines("input-example.txt")

	got := findEncryptionWeakness(values, 127)
	if got != 62 {
		t.Errorf("wanted 62, got %d", got)
	}
}
