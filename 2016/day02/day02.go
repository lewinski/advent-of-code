package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	simpleKeypad := util.IntGrid2{
		util.Point2{-1, 1}:  '1',
		util.Point2{0, 1}:   '2',
		util.Point2{1, 1}:   '3',
		util.Point2{-1, 0}:  '4',
		util.Point2{0, 0}:   '5',
		util.Point2{1, 0}:   '6',
		util.Point2{-1, -1}: '7',
		util.Point2{0, -1}:  '8',
		util.Point2{1, -1}:  '9',
	}
	fmt.Println("part1:", decode(lines, simpleKeypad))

	advancedKeypad := util.IntGrid2{
		util.Point2{0, 2}:   '1',
		util.Point2{-1, 1}:  '2',
		util.Point2{0, 1}:   '3',
		util.Point2{1, 1}:   '4',
		util.Point2{-2, 0}:  '5',
		util.Point2{-1, 0}:  '6',
		util.Point2{0, 0}:   '7',
		util.Point2{1, 0}:   '8',
		util.Point2{2, 0}:   '9',
		util.Point2{-1, -1}: 'A',
		util.Point2{0, -1}:  'B',
		util.Point2{1, -1}:  'C',
		util.Point2{0, -2}:  'D',
	}
	fmt.Println("part2:", decode(lines, advancedKeypad))
}

func decode(lines []string, keypad util.IntGrid2) string {
	directions := map[rune]util.Point2{
		'U': {0, 1},
		'D': {0, -1},
		'L': {-1, 0},
		'R': {1, 0},
	}
	pos := util.Origin2()
	code := make([]rune, 0, len(lines))

	for _, line := range lines {
		for _, c := range line {
			offset := directions[c]
			next := pos.Offset(offset)
			if keypad.Contains(next) {
				pos = next
			}
		}
		code = append(code, (rune)(keypad.Get(pos)))
	}

	return string(code)
}
