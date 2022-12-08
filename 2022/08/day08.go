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

	fmt.Println("part1:", part1(g))
	fmt.Println("part2:", part2(g))
}

func part1(g util.IntGrid2) int {
	dirs := util.Origin2().Touching()

	visible := 0
	g.Each(func(p util.Point2, v int) {
		for _, dir := range dirs {
			// move in direction until we hit a tree that is too tall or fall off the edge
			cur := p.Offset(dir)
			for g.Contains(cur) {
				if g.Get(cur) >= v {
					break
				}
				cur = cur.Offset(dir)
			}
			// if we left the edge, then the tree is visible from outside
			if !g.Contains(cur) {
				visible++
				return
			}
		}
	})
	return visible
}

func part2(g util.IntGrid2) int {
	dirs := util.Origin2().Touching()

	best := 0
	g.Each(func(p util.Point2, x int) {
		score := 1
		for _, dir := range dirs {
			trees := 0

			// move in direction, counting trees that are shorter than starting position
			cur := p.Offset(dir)
			for g.Contains(cur) {
				trees++
				if g.Get(cur) >= x {
					break
				}
				cur = cur.Offset(dir)
			}

			// accumulate score
			score *= trees
		}

		if score > best {
			best = score
		}
	})
	return best
}
