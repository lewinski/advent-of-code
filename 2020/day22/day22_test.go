package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Records("input-example.txt")

	deck1 := parseDeck(input[0])
	deck2 := parseDeck(input[1])

	deck := combatGame(deck1, deck2)

	got := deck.score()
	if got != 306 {
		t.Errorf("wanted score 306, got %d", got)
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Records("input-example.txt")

	deck1 := parseDeck(input[0])
	deck2 := parseDeck(input[1])

	deck, _ := recursiveCombatGame(deck1, deck2)

	got := deck.score()
	if got != 291 {
		t.Errorf("wanted score 291, got %d", got)
	}
}

func TestInfiniteLoopPart2(t *testing.T) {
	input := util.Records("input-inf.txt")

	deck1 := parseDeck(input[0])
	deck2 := parseDeck(input[1])

	deck, _ := recursiveCombatGame(deck1, deck2)

	got := deck.score()
	if got != 105 {
		t.Errorf("wanted score 105, got %d", got)
	}
}
