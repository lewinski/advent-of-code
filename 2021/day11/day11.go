package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	fmt.Println("part1:", part1(parse(lines)))
	fmt.Println("part2:", part2(parse(lines)))
}

func parse(lines []string) util.IntGrid2 {
	octos := util.IntGrid2{}
	for x, line := range lines {
		for y, c := range line {
			octos.SetCoords(x, y, util.MustAtoi(string(c)))
		}
	}
	return octos
}

func step(grid util.IntGrid2) (util.IntGrid2, int) {
	flashed := util.IntGrid2{}

	grid.Each(func(p util.Point2, x int) {
		grid.Set(p, x+1)
	})

	for {
		again := false
		grid.Each(func(p util.Point2, x int) {
			if grid.Get(p) > 9 && !flashed.Contains(p) {
				flashed.Set(p, 1)
				again = true
				for _, p2 := range p.Around() {
					if grid.Contains(p2) {
						grid.Set(p2, grid.Get(p2)+1)
					}
				}
			}
		})
		if !again {
			break
		}
	}

	grid.Each(func(p util.Point2, x int) {
		if x > 9 {
			grid.Set(p, 0)
		}
	})

	return grid, len(flashed)
}

func part1(grid util.IntGrid2) int {
	total := 0
	for i := 0; i < 100; i++ {
		var flashes int
		grid, flashes = step(grid)
		total += flashes
	}
	return total
}

func part2(grid util.IntGrid2) int {
	for i := 1; ; i++ {
		var flashes int
		grid, flashes = step(grid)
		if flashes == 100 {
			return i
		}
	}
}
