package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Lines("input-example.txt")
	f := parseInput(input)
	for {
		g := f.iter1()
		if f.String() == g.String() {
			break
		}
		f = g
	}

	count := f.occupiedSeats()
	if count != 37 {
		t.Errorf("wanted 37 occupied seats, got %d", count)
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Lines("input-example.txt")
	f := parseInput(input)

	for {
		g := f.iter2()
		if f.String() == g.String() {
			break
		}
		f = g
	}

	count := f.occupiedSeats()
	if count != 26 {
		t.Errorf("wanted 26 occupied seats, got %d", count)
	}
}
