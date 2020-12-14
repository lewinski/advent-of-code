package main

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestExamplePart1(t *testing.T) {
	input := util.Lines("input-example.txt")
	timestamp := util.MustAtoi(input[0])
	schedule := input[1]

	gotID, gotTimestamp := earliestBus(timestamp, schedule)
	if gotID != 59 {
		t.Errorf("wanted bus 59, got %d", gotID)
	}
	if gotTimestamp-timestamp != 5 {
		t.Errorf("wanted next departure in 5 minutes, got %d", gotTimestamp-timestamp)
	}
}

func TestExamplePart2(t *testing.T) {
	tests := []struct {
		schedule string
		want     int
	}{
		{"17,x,13,19", 3417},
		{"7,13,x,x,59,x,31,19", 1068781},
		{"67,7,59,61", 754018},
		{"67,x,7,59,61", 779210},
		{"67,7,x,59,61", 1261476},
		{"1789,37,47,1889", 1202161486},
	}

	for i, tt := range tests {
		got := contestTimestamp(tt.schedule)
		if got != tt.want {
			t.Errorf("test %d: wanted timestamp to be %d, got %d", i, tt.want, got)
		}
	}
}
