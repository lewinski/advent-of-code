package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", moveRope(lines, 2))
	fmt.Println("part1:", moveRope(lines, 10))
}

func moveRope(lines []string, n int) int {
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
				// where are we in relation to knot in front of us
				dx := knots[j-1][0] - knots[j][0]
				dy := knots[j-1][1] - knots[j][1]

				if util.IAbs(dx) <= 1 && util.IAbs(dy) <= 1 {
					// already touching
					continue
				}

				// move closer
				knots[j] = knots[j].OffsetCoords(util.ISign(dx), util.ISign(dy))
			}

			seen.Set(knots[n-1], 1)
		}
	}

	return len(seen)
}
