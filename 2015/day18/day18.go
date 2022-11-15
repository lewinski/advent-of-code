package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")
	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func parseGrid(lines []string) util.IntGrid2 {
	grid := util.IntGrid2{}

	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				grid.SetCoords(x, y, 1)
			}
		}
	}

	return grid
}

func iter(grid util.IntGrid2, size int) util.IntGrid2 {
	next := util.IntGrid2{}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			count := 0
			p := util.Point2{x, y}
			for _, p := range p.Around() {
				count += grid.Get(p)
			}
			if grid.GetCoords(x, y) == 1 && (count == 2 || count == 3) {
				next.SetCoords(x, y, 1)
			}
			if grid.GetCoords(x, y) == 0 && count == 3 {
				next.SetCoords(x, y, 1)
			}
		}
	}
	return next
}

func part1(lines []string) int {
	grid := parseGrid(lines)

	for t := 0; t < 100; t++ {
		grid = iter(grid, 100)
	}

	total := 0
	grid.Each(func(p util.Point2, v int) {
		total += v
	})
	return total
}

func part2(lines []string) int {
	grid := parseGrid(lines)

	for t := 0; t < 100; t++ {
		grid = iter(grid, 100)
		grid.SetCoords(0, 0, 1)
		grid.SetCoords(99, 0, 1)
		grid.SetCoords(0, 99, 1)
		grid.SetCoords(99, 99, 1)
	}

	total := 0
	grid.Each(func(p util.Point2, v int) {
		total += v
	})
	return total
}
