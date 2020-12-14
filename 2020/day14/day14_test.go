package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Lines("input-example.txt")

	memory := bootloaderVersion1(input)
	got := sumMemory(memory)

	var want int = 165
	if want != got {
		t.Errorf("wanted sum to be %d, got %d", want, got)
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Lines("input-example2.txt")

	memory := bootloaderVersion2(input)
	got := sumMemory(memory)

	var want int = 208
	if want != got {
		t.Errorf("wanted sum to be %d, got %d", want, got)
	}
}
