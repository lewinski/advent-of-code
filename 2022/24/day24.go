package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

const (
	OPEN int = 0
	WALL     = 1 << iota
	UP
	DOWN
	LEFT
	RIGHT
)

func main() {
	lines := util.Lines("input.txt")

	g := util.IntGrid2{}

	for y, line := range lines {
		for x, c := range line {
			switch c {
			case '.':
				g.SetCoords(x, y, OPEN)
			case '#':
				g.SetCoords(x, y, WALL)
			case '^':
				g.SetCoords(x, y, UP)
			case 'v':
				g.SetCoords(x, y, DOWN)
			case '<':
				g.SetCoords(x, y, LEFT)
			case '>':
				g.SetCoords(x, y, RIGHT)
			}
		}
	}

	bmin, bmax := g.Bounds()
	entrance := util.Point2{bmin[0] + 1, bmin[1]}
	exit := util.Point2{bmax[0] - 1, bmax[1]}

	part1, g := minStepsToGoal(g, entrance, exit)
	fmt.Println("part1:", part1)

	back, g := minStepsToGoal(g, exit, entrance)
	again, g := minStepsToGoal(g, entrance, exit)
	fmt.Println("part2:", part1+back+again)
}

func minStepsToGoal(grid util.IntGrid2, start, goal util.Point2) (int, util.IntGrid2) {
	bmin, bmax := grid.Bounds()

	pos := map[util.Point2]bool{}
	pos[start] = true

	steps := 0
	for {
		steps++

		next := tick(grid, bmin, bmax)
		nextPos := map[util.Point2]bool{}

		for p := range pos {
			if next.Get(p) == OPEN {
				nextPos[p] = true
			}
			for _, q := range p.Touching() {
				if next.Contains(q) && next.Get(q) == OPEN {
					nextPos[q] = true
				}
			}
		}

		grid = next
		pos = nextPos

		if pos[goal] {
			return steps, grid
		}
	}
}

func tick(g util.IntGrid2, bmin, bmax util.Point2) util.IntGrid2 {
	next := util.IntGrid2{}

	g.Each(func(p util.Point2, v int) {
		if !next.Contains(p) {
			next.Set(p, OPEN)
		}
		if v == WALL {
			next.Set(p, WALL)
			return
		}
		if v&UP != 0 {
			q := p.OffsetCoords(0, -1)
			if q[1] == bmin[1] {
				q[1] = bmax[1] - 1
			}
			next.Set(q, next.Get(q)|UP)
		}
		if v&DOWN != 0 {
			q := p.OffsetCoords(0, 1)
			if q[1] == bmax[1] {
				q[1] = bmin[1] + 1
			}
			next.Set(q, next.Get(q)|DOWN)
		}
		if v&LEFT != 0 {
			q := p.OffsetCoords(-1, 0)
			if q[0] == bmin[0] {
				q[0] = bmax[0] - 1
			}
			next.Set(q, next.Get(q)|LEFT)
		}
		if v&RIGHT != 0 {
			q := p.OffsetCoords(1, 0)
			if q[0] == bmax[0] {
				q[0] = bmin[0] + 1
			}
			next.Set(q, next.Get(q)|RIGHT)
		}
	})

	return next
}

func printGrid(g util.IntGrid2, p map[util.Point2]bool) {
	bmin, bmax := g.Bounds()
	for y := bmin[1]; y <= bmax[1]; y++ {
		var sb strings.Builder
		for x := bmin[0]; x <= bmax[0]; x++ {
			if p[util.Point2{x, y}] {
				sb.WriteString("E")
				continue
			}
			switch g.GetCoords(x, y) {
			case OPEN:
				sb.WriteString(".")
			case WALL:
				sb.WriteString("#")
			case UP:
				sb.WriteString("^")
			case DOWN:
				sb.WriteString("v")
			case LEFT:
				sb.WriteString("<")
			case RIGHT:
				sb.WriteString(">")
			default:
				sb.WriteString("?")
			}
		}
		fmt.Println(sb.String())
	}
}
