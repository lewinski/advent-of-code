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

func drawLine(g *util.IntGrid2, x1, y1, x2, y2 int) {
	xdir := 0
	if x1 < x2 {
		xdir = 1
	} else if x1 > x2 {
		xdir = -1
	}

	ydir := 0
	if y1 < y2 {
		ydir = 1
	} else if y1 > y2 {
		ydir = -1
	}

	for x, y := x1, y1; ; x, y = x+xdir, y+ydir {
		g.Set(util.Point2{x, y}, g.Get(util.Point2{x, y})+1)
		if x == x2 && y == y2 {
			break
		}
	}
}

func part1(lines []string) int {
	g := util.IntGrid2{}

	for _, line := range lines {
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		if x1 == x2 || y1 == y2 {
			drawLine(&g, x1, y1, x2, y2)
		}
	}

	sum := 0
	g.Each(func(p util.Point2, x int) {
		if x >= 2 {
			sum++
		}
	})

	return sum
}

func part2(lines []string) int {
	g := util.IntGrid2{}

	for _, line := range lines {
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		drawLine(&g, x1, y1, x2, y2)
	}

	sum := 0
	g.Each(func(p util.Point2, x int) {
		if x >= 2 {
			sum++
		}
	})

	return sum
}
