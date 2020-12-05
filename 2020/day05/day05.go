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

func decodeSeat(pass string) (row, col int) {
	rowMin, rowMax := 0, 127
	for _, x := range pass[:7] {
		if x == 'F' {
			rowMax -= (rowMax + 1 - rowMin) / 2
		} else if x == 'B' {
			rowMin += (rowMax + 1 - rowMin) / 2
		}
	}
	row = rowMin

	colMin, colMax := 0, 7
	for _, x := range pass[7:] {
		if x == 'L' {
			colMax -= (colMax + 1 - colMin) / 2
		} else if x == 'R' {
			colMin += (colMax + 1 - colMin) / 2
		}
	}
	col = colMin

	return
}

func seatID(row, col int) int {
	return row*8 + col
}
