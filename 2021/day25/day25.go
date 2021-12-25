package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	g := parseGrid(lines)
	size := util.Point2{len(lines[0]), len(lines)}

	fmt.Println("part1:", part1(g, size))
}

func parseGrid(lines []string) util.IntGrid2 {
	g := util.IntGrid2{}

	for y, line := range lines {
		for x, c := range line {
			switch c {
			case '.':
				// g.SetCoords(x, y, 0)
			case '>':
				g.SetCoords(x, y, 1)
			case 'v':
				g.SetCoords(x, y, 2)
			default:
				panic("oh no")
			}
		}
	}

	return g
}

func printGrid(g util.IntGrid2, size util.Point2) {
	for y := 0; y < size[1]; y++ {
		for x := 0; x < size[0]; x++ {
			switch g.GetCoords(x, y) {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print(">")
			case 2:
				fmt.Print("v")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("")
}

func part1(g util.IntGrid2, size util.Point2) int {
	steps := 0

	for {
		steps++

		moves := 0

		g2 := util.IntGrid2{}
		g.Each(func(p util.Point2, x int) {
			p2 := util.Point2{(p[0] + 1) % size[0], p[1]}
			if x == 1 && g.Get(p2) == 0 {
				moves++
				g2.Set(p2, x)
			} else {
				g2.Set(p, x)
			}
		})
		g = g2

		g2 = util.IntGrid2{}
		g.Each(func(p util.Point2, x int) {
			p2 := util.Point2{p[0], (p[1] + 1) % size[1]}
			if x == 2 && g.Get(p2) == 0 {
				moves++
				g2.Set(p2, x)
			} else {
				g2.Set(p, x)
			}
		})
		g = g2

		if moves == 0 {
			break
		}
	}

	return steps
}
