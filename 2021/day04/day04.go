package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

type board [5][5]struct {
	number int
	marked bool
}

func (b *board) mark(n int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j].number == n {
				b[i][j].marked = true
			}
		}
	}
}

func (b *board) won() bool {
	for i := 0; i < 5; i++ {
		if b[0][i].marked && b[1][i].marked && b[2][i].marked && b[3][i].marked && b[4][i].marked {
			return true
		}
		if b[i][0].marked && b[i][1].marked && b[i][2].marked && b[i][3].marked && b[i][4].marked {
			return true
		}
	}
	return false
}

func (b *board) score() int {
	s := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b[i][j].marked {
				s += b[i][j].number
			}
		}
	}
	return s
}

func parseCalls(calls string) []int {
	nums := []int{}
	for _, x := range strings.Split(calls, ",") {
		nums = append(nums, util.MustAtoi(x))
	}
	return nums
}

func parseBoards(records []string) []board {
	boards := []board{}

	for _, record := range records {
		b := board{}

		for i, l := range strings.SplitN(record, "\n", 5) {
			for j, x := range strings.Fields(l) {
				b[i][j].number = util.MustAtoi(x)
			}
		}

		boards = append(boards, b)
	}

	return boards
}

func main() {
	records := util.Records("input.txt")

	fmt.Println("part1:", part1(records))
	fmt.Println("part2:", part2(records))
}

func part1(records []string) int {
	calls := parseCalls(records[0])
	boards := parseBoards(records[1:])

	for _, x := range calls {
		for i := range boards {
			boards[i].mark(x)
			if boards[i].won() {
				return boards[i].score() * x
			}
		}
	}
	panic("failed")
}

func part2(records []string) int {
	calls := parseCalls(records[0])
	boards := parseBoards(records[1:])
	wins := 0

	for _, x := range calls {
		for i := range boards {
			if boards[i].won() {
				// already won, skip marking
				continue
			}
			boards[i].mark(x)
			if boards[i].won() {
				wins++
				if wins == len(boards) {
					return boards[i].score() * x
				}
			}
		}
	}
	panic("failed")
}
