package main

import (
	"fmt"
	"sort"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	g := util.IntGrid2{}
	for x, line := range lines {
		for y, c := range line {
			g.Set(util.Point2{x, y}, util.MustAtoi(string(c)))
		}
	}

	fmt.Println("part1:", part1(g))
	fmt.Println("part2:", part2(g))
}

func part1(grid util.IntGrid2) int {
	risk := 0
	grid.Each(func(p util.Point2, val int) {
		for _, p2 := range p.Touching() {
			if grid.Contains(p2) && grid.Get(p2) <= val {
				return
			}
		}
		risk += val + 1
	})
	return risk
}

func flood(grid *util.IntGrid2, p util.Point2, count *int) {
	if grid.Contains(p) && grid.Get(p) != 9 {
		*count++
		grid.Set(p, 9)
		for _, p2 := range p.Touching() {
			flood(grid, p2, count)
		}
	}
}

func part2(grid util.IntGrid2) int {
	basins := []int{}

	grid.Each(func(p util.Point2, val int) {
		for _, p2 := range p.Touching() {
			if grid.Contains(p2) && grid.Get(p2) <= val {
				return
			}
		}

		size := 0
		flood(&grid, p, &size)
		basins = append(basins, size)
	})

	sort.Ints(basins)

	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}
