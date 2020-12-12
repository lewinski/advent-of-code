package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Lines("input-example.txt")
	ferry := parseInput(input)
	for {
		gerry := ferry.iter1()
		if ferry.String() == gerry.String() {
			break
		}
		ferry = gerry
	}

	count := ferry.occupiedSeats()
	if count != 37 {
		t.Errorf("wanted 37 occupied seats, got %d", count)
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Lines("input-example.txt")
	ferry := parseInput(input)
	for {
		gerry := ferry.iter2()
		if ferry.String() == gerry.String() {
			break
		}
		ferry = gerry
	}

	count := ferry.occupiedSeats()
	if count != 26 {
		t.Errorf("wanted 26 occupied seats, got %d", count)
	}
}
