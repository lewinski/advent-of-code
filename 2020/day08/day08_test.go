package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	program := parseProgram(util.Lines("input-example.txt"))
	got, booted := runProgram(program)
	want := 5
	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
	if booted {
		t.Error("booted successfully, but should have looped")
	}
}

func TestExamplePart2(t *testing.T) {
	program := parseProgram(util.Lines("input-example.txt"))
	got, booted := fixProgram(program)
	want := 8
	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
	if !booted {
		t.Error("should have booted successfully, but didn't")
	}
}
