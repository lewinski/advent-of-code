package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Records("input-example.txt")

	fieldMap := parseClasses(input[0])
	otherTickets := parseTickets(input[2])

	validTickets, errorRate := validTickets(otherTickets, fieldMap)

	if errorRate != 71 {
		t.Errorf("wanted error rate = 71, got %d", errorRate)
	}
	if len(validTickets) != 1 {
		t.Errorf("wanted 1 valid ticket, got %d", len(validTickets))
	}
}

func TestExamplePart2(t *testing.T) {
	input := util.Records("input-example2.txt")

	fieldMap := parseClasses(input[0])
	otherTickets := parseTickets(input[2])

	validTickets, _ := validTickets(otherTickets, fieldMap)
	fieldPos := solveFieldMapping(validTickets, fieldMap)

	if fieldPos["row"] != 0 {
		t.Errorf("wanted row field position to be 0, got %d", fieldPos["row"])
	}
	if fieldPos["class"] != 1 {
		t.Errorf("wanted class field position to be 1, got %d", fieldPos["class"])
	}
	if fieldPos["seat"] != 2 {
		t.Errorf("wanted seat field position to be 2, got %d", fieldPos["seat"])
	}
}
