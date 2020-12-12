package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Lines("input-example.txt")
	got := p1nav(input)
	if got.x != 17 || got.y != -8 {
		t.Errorf("wanted position = {17, -8}, got %v", got)
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Lines("input-example.txt")
	got := p2nav(input)
	if got.x != 214 || got.y != -72 {
		t.Errorf("wanted position = {214, -72}, got %v", got)
	}
}
