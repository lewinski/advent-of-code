package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

const (
	CLEAN int = iota
	WEAKENED
	INFECTED
	FLAGGED
)

func parseGrid(filename string) util.IntGrid2 {
	g := util.IntGrid2{}
	for y, line := range util.Lines(filename) {
		for x, c := range line {
			offset := len(line) / 2
			if c == '#' {
				g.SetCoords(x-offset, offset-y, INFECTED)
			}
		}
	}
	return g
}

func main() {
	fmt.Println("part1:", part1(parseGrid("input.txt")))
	fmt.Println("part1:", part2(parseGrid("input.txt")))
}

func part1(g util.IntGrid2) int {
	pos := util.Origin2()
	dir := util.Point2{0, 1}

	infections := 0
	for i := 0; i < 10000; i++ {
		switch g.Get(pos) {
		case CLEAN:
			dir = util.Point2{-dir[1], dir[0]}
			g.Set(pos, INFECTED)
			infections++
		case INFECTED:
			dir = util.Point2{dir[1], -dir[0]}
			g.Set(pos, CLEAN)
		}
		pos = pos.Offset(dir)
	}

	return infections
}

func part2(g util.IntGrid2) int {
	pos := util.Origin2()
	dir := util.Point2{0, 1}

	infections := 0
	for i := 0; i < 10000000; i++ {
		switch g.Get(pos) {
		case CLEAN:
			dir = util.Point2{-dir[1], dir[0]}
			g.Set(pos, WEAKENED)
		case WEAKENED:
			g.Set(pos, INFECTED)
			infections++
		case INFECTED:
			dir = util.Point2{dir[1], -dir[0]}
			g.Set(pos, FLAGGED)
		case FLAGGED:
			dir = util.Point2{-dir[0], -dir[1]}
			g.Set(pos, CLEAN)
		}
		pos = pos.Offset(dir)
	}

	return infections
}
