package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	program := parseProgram(util.Lines("input-example.txt"))
	got, looped := runProgram(program)
	want := 5
	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
	if !looped {
		t.Error("wanted program looped, got booted instead")
	}
}

func TestExamplePart2(t *testing.T) {
	program := parseProgram(util.Lines("input-example.txt"))
	got, looped := fixProgram(program)
	want := 8
	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
	if looped {
		t.Error("wanted program booted, got looped instead")
	}
}
