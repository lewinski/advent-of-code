package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", part1(lines))
	fmt.Println("part1:", part2(lines, 10))
}

func part1(lines []string) int {
	directions := map[byte]util.Point2{
		'R': {1, 0},
		'L': {-1, 0},
		'U': {0, 1},
		'D': {0, -1},
	}

	seen := util.IntGrid2{}
	head := util.Origin2()
	tail := util.Origin2()

	seen.Set(tail, 1)

	for _, line := range lines {
		dir := directions[line[0]]
		count := util.MustAtoi(line[2:])

		for i := 0; i < count; i++ {
			head = head.Offset(dir)

			dx := util.IAbs(head[0] - tail[0])
			dy := util.IAbs(head[1] - tail[1])
			if dx <= 1 && dy <= 1 {
				continue
			} else {
				// tail will always move to the place the head was
				tail = head.Offset(dir.Scale(-1))
			}

			seen.Set(tail, 1)
		}
	}

	return len(seen)
}

func part2(lines []string, n int) int {
	directions := map[byte]util.Point2{
		'R': {1, 0},
		'L': {-1, 0},
		'U': {0, 1},
		'D': {0, -1},
	}

	seen := util.IntGrid2{}
	knots := make([]util.Point2, n)
	seen.Set(knots[n-1], 1)

	for _, line := range lines {
		dir := directions[line[0]]
		count := util.MustAtoi(line[2:])

		for i := 0; i < count; i++ {
			// move head
			knots[0] = knots[0].Offset(dir)

			// move knots behind tail
			for j := 1; j < n; j++ {
				dx := knots[j][0] - knots[j-1][0]
				dy := knots[j][1] - knots[j-1][1]

				if util.IAbs(dx) <= 1 && util.IAbs(dy) <= 1 {
					// already touching
					continue
				} else if util.IAbs(dx) == 2 && dy == 0 {
					// close gap directly left or right
					knots[j] = knots[j].OffsetCoords(-dx/util.IAbs(dx), 0)
				} else if dx == 0 && util.IAbs(dy) == 2 {
					// close gap directly above or below
					knots[j] = knots[j].OffsetCoords(0, -dy/util.IAbs(dy))
				} else {
					// move diagonally
					knots[j] = knots[j].OffsetCoords(-dx/util.IAbs(dx), -dy/util.IAbs(dy))
				}
			}

			seen.Set(knots[n-1], 1)
		}
	}

	return len(seen)
}
