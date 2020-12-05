package main

import (
	"fmt"
	"math"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")

	minID, maxID := math.MaxInt16, math.MinInt16
	for _, line := range input {
		id := seatID(decodeSeat(line))
		if id < minID {
			minID = id
		}
		if id > maxID {
			maxID = id
		}
	}
	fmt.Println("part1:", maxID)

	seats := map[int]bool{}
	for i := minID; i < maxID; i++ {
		seats[i] = true
	}
	for _, line := range input {
		id := seatID(decodeSeat(line))
		delete(seats, id)
	}
	if len(seats) != 1 {
		panic("oh no")
	}
	for i := range seats {
		fmt.Println("part2:", i)
	}
}

func bsp(min, max int, rule string, lower, upper rune) int {
	for _, x := range rule {
		if x == lower {
			max -= (max - min + 1) / 2
		} else if x == upper {
			min += (max - min + 1) / 2
		}
	}
	return min
}

func decodeSeat(pass string) (row, col int) {
	row = bsp(0, 127, pass[:7], 'F', 'B')
	col = bsp(0, 7, pass[7:], 'L', 'R')
	return
}

func seatID(row, col int) int {
	return row*8 + col
}
