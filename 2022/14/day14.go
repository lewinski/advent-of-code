package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")
	grid, lowest := parseGrid(lines)

	part1, part2 := simulate(grid, lowest)
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func simulate(grid util.IntGrid2, lowest int) (part1, part2 int) {
	moves := []util.Point2{{0, 1}, {-1, 1}, {1, 1}}

	for {
		sand := util.Point2{500, 0}

	drop:
		for {
			// part 1: entering freefall
			if part1 == 0 && sand[1] > lowest {
				part1 = part2
			}

			// check available moves
			for _, move := range moves {
				next := sand.Offset(move)
				// part 2: floor is at lowest + 2
				if next[1] != lowest+2 && grid.Get(next) == 0 {
					sand = next
					continue drop
				}
			}

			// no more moves, so drop a sand here
			grid.Set(sand, 'o')
			part2++

			// and prepare to start another particle
			break
		}

		// sand didn't fall, we should exit
		if sand[1] == 0 {
			return
		}
	}
}

func parseGrid(lines []string) (util.IntGrid2, int) {
	lowest := 0

	grid := util.IntGrid2{}
	for _, line := range lines {
		points := []util.Point2{}
		for _, p := range strings.Split(line, " -> ") {
			f := strings.SplitN(p, ",", 2)
			points = append(points, util.Point2{util.MustAtoi(f[0]), util.MustAtoi(f[1])})
		}

		for i := 1; i < len(points); i++ {
			delta := util.Point2{
				util.ISign(points[i][0] - points[i-1][0]),
				util.ISign(points[i][1] - points[i-1][1]),
			}

			for pos := points[i-1]; ; pos = pos.Offset(delta) {
				lowest = util.IMax(lowest, pos[1])
				grid[pos] = '#'
				if pos == points[i] {
					break
				}
			}
		}
	}

	return grid, lowest
}

func printGrid(grid util.IntGrid2) {
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt

	grid.Each(func(p util.Point2, v int) {
		if p[0] < minX {
			minX = p[0]
		}
		if p[0] > maxX {
			maxX = p[0]
		}
		if p[1] < minY {
			minY = p[1]
		}
		if p[1] > maxY {
			maxY = p[1]
		}
	})

	for y := 0; y <= maxY; y++ {
		var sb strings.Builder
		for x := minX; x <= maxX; x++ {
			if grid.ContainsCoords(x, y) {
				sb.WriteByte(byte(grid.GetCoords(x, y)))
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
		fmt.Print(sb.String())
	}
}
