package main

import "fmt"

// Disc #1 has 7 positions; at time=0, it is at position 0.
// Disc #2 has 13 positions; at time=0, it is at position 0.
// Disc #3 has 3 positions; at time=0, it is at position 2.
// Disc #4 has 5 positions; at time=0, it is at position 2.
// Disc #5 has 17 positions; at time=0, it is at position 0.
// Disc #6 has 19 positions; at time=0, it is at position 7.

func main() {
	discs := []disc{
		{7, 0},
		{13, 0},
		{3, 2},
		{5, 2},
		{17, 0},
		{19, 7},
	}
	fmt.Println("part1:", solve(discs))

	discs = []disc{
		{7, 0},
		{13, 0},
		{3, 2},
		{5, 2},
		{17, 0},
		{19, 7},
		{11, 0},
	}
	fmt.Println("part2:", solve(discs))
}

type disc struct {
	positions int
	current   int
}

func solve(discs []disc) int {
	for j := 0; j < len(discs); j++ {
		discs[j].current = (discs[j].current + j + 1) % discs[j].positions
	}

	for i := 1; ; i++ {
		done := true
		for j := 0; j < len(discs); j++ {
			discs[j].current = (discs[j].current + 1) % discs[j].positions
			if discs[j].current != 0 {
				done = false
			}
		}
		if done {
			return i
		}
	}
}
