package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	input := "389125467"
	want := []string{
		"389125467", // start
		"289154673", // turn 1
		"546789132", // turn 2
		"891346725", // turn 3
		"467913258", // turn 4
		"136792584", // turn 5
		"936725841", // turn 6
		"258367419", // turn 7
		"674158392", // turn 8
		"574183926", // turn 9
		"837419265", // turn 10
	}

	current, cups := stringToCup(input)
	got := cupToString(current)
	if got != want[0] {
		t.Errorf("turn %d: wanted %s, got %s", 0, want[0], got)
	}

	for turn := 1; turn < len(want); turn++ {
		current = move(current, cups)
		got = cupToString(current)
		if got != want[turn] {
			t.Errorf("turn %d: wanted %s, got %s", turn, want[turn], got)
		}
	}

	oneWant := "92658374"
	oneGot := cupToString(&cups[0])[1:]
	if oneGot != oneWant {
		t.Errorf("turn 10: wanted %s after 1, got %s", oneWant, oneGot)
	}

	current, cups = stringToCup(input)
	for i := 0; i < 100; i++ {
		current = move(current, cups)
	}

	oneWant = "67384529"
	oneGot = cupToString(&cups[0])[1:]
	if oneGot != oneWant {
		t.Errorf("turn 100: wanted %s after 1, got %s", oneWant, oneGot)
	}
}

func TestExamplePart2(t *testing.T) {
	input := "389125467"
	current, cups := stringToCup2(input)
	for i := 0; i < 10000000; i++ {
		current = move(current, cups)
	}

	one := &cups[0]
	got := one.next.label * one.next.next.label
	want := 149245887792

	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
}
