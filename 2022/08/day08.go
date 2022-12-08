package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	g := util.IntGrid2{}
	for y, line := range lines {
		for x, c := range line {
			g.SetCoords(x, y, int(c-'0'))
		}
	}

	h := len(lines)
	w := len(lines[0])
	fmt.Println("part1:", part1(g, h, w))
	fmt.Println("part2:", part2(g))
}

func part1(g util.IntGrid2, h, w int) int {
	visible := 0
	g.Each(func(p util.Point2, v int) {
		if p[0] == 0 || p[0] == w-1 || p[1] == 0 || p[1] == h-1 {
			// everything on edge is visible
			visible++
		} else {
			ok := true
			for x := 0; x < p[0]; x++ {
				if g.GetCoords(x, p[1]) >= v {
					ok = false
				}
			}
			if ok {
				visible++
				return
			}

			ok = true
			for x := w - 1; x > p[0]; x-- {
				if g.GetCoords(x, p[1]) >= v {
					ok = false
				}
			}
			if ok {
				visible++
				return
			}

			ok = true
			for y := 0; y < p[1]; y++ {
				if g.GetCoords(p[0], y) >= v {
					ok = false
				}
			}
			if ok {
				visible++
				return
			}

			ok = true
			for y := h - 1; y > p[1]; y-- {
				if g.GetCoords(p[0], y) >= v {
					ok = false
				}
			}
			if ok {
				visible++
				return
			}
		}
	})
	return visible
}

func part2(g util.IntGrid2) int {
	best := 0

	g.Each(func(p util.Point2, x int) {
		dirs := []util.Point2{
			{1, 0},
			{0, 1},
			{-1, 0},
			{0, -1},
		}

		score := 1
		for _, dir := range dirs {
			c := p.Offset(dir)
			see := 0
			for g.Contains(c) {
				see++
				if g.Get(c) >= x {
					break
				}
				c = c.Offset(dir)
			}
			score *= see
		}

		if score > best {
			best = score
		}
	})

	return best
}
