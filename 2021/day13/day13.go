package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	records := util.Records("input.txt")

	points := util.IntGrid2{}
	for _, line := range strings.Split(records[0], "\n") {
		fields := strings.Split(line, ",")
		points.SetCoords(util.MustAtoi(fields[0]), util.MustAtoi(fields[1]), 1)
	}
	instructions := strings.Split(records[1], "\n")

	fmt.Println("part1:", part1(points, instructions))
	fmt.Printf("part2:\n%s", part2(points, instructions))
}

func foldUp(g util.IntGrid2, y int) util.IntGrid2 {
	h := util.IntGrid2{}
	g.Each(func(p util.Point2, v int) {
		if p[1] < y {
			h.Set(p, v)
		} else {
			h.SetCoords(p[0], 2*y-p[1], v)
		}
	})
	return h
}

func foldLeft(g util.IntGrid2, x int) util.IntGrid2 {
	h := util.IntGrid2{}
	g.Each(func(p util.Point2, v int) {
		if p[0] < x {
			h.Set(p, v)
		} else {
			h.SetCoords(2*x-p[0], p[1], v)
		}
	})
	return h
}

func fold(g util.IntGrid2, instruction string) util.IntGrid2 {
	f := strings.Split(instruction, "=")
	switch f[0] {
	case "fold along y":
		return foldUp(g, util.MustAtoi(f[1]))
	case "fold along x":
		return foldLeft(g, util.MustAtoi(f[1]))
	default:
		panic("unknown instruction")
	}
}

func part1(g util.IntGrid2, instructions []string) int {
	return len(fold(g, instructions[0]))
}

func part2(g util.IntGrid2, instructions []string) string {
	for i := range instructions {
		g = fold(g, instructions[i])
	}
	return formatGrid(g)
}

func formatGrid(g util.IntGrid2) string {
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	g.Each(func(p util.Point2, x int) {
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

	var sb strings.Builder
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if g.ContainsCoords(x, y) {
				sb.WriteRune('\u2588')
			} else {
				sb.WriteRune(' ')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
