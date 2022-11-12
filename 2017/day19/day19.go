package main

import (
	"errors"
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")
	grid := parseGrid(input)

	fmt.Println("part1:", part1(grid))
	fmt.Println("part2:", part2(grid))
}

func parseGrid(lines []string) util.IntGrid2 {
	g := util.IntGrid2{}

	for y, line := range lines {
		for x, c := range line {
			if c != ' ' {
				g.SetCoords(x, y, int(c))
			}
		}
	}

	return g
}

func findStart(grid util.IntGrid2) util.Point2 {
	var start util.Point2
	grid.Each(func(p util.Point2, x int) {
		if x == int('|') && p[1] == 0 {
			start = p
		}
	})
	return start
}

func gridMove(grid util.IntGrid2, pos, dir util.Point2) (util.Point2, util.Point2, error) {
	var next util.Point2

	try := []util.Point2{dir, {dir[1], -dir[0]}, {-dir[1], dir[0]}}
	for _, p := range try {
		if grid.Contains(pos.Offset(p)) {
			dir = p
			next = pos.Offset(p)
			break
		}
	}

	if next == util.Origin2() {
		return pos, dir, errors.New("no move available")
	}

	return next, dir, nil
}

func part1(grid util.IntGrid2) string {
	visited := []rune{}

	pos := findStart(grid)
	dir := util.Point2{0, 1}

	for {
		var err error
		pos, dir, err = gridMove(grid, pos, dir)
		if err != nil {
			break
		}

		v := rune(grid.Get(pos))
		switch v {
		case '|': // ignore
		case '+': // ignore
		case '-': // ignore
		default:
			visited = append(visited, v)
		}
	}

	return string(visited)
}

func part2(grid util.IntGrid2) int {
	steps := 1

	pos := findStart(grid)
	dir := util.Point2{0, 1}

	for {
		var err error
		pos, dir, err = gridMove(grid, pos, dir)
		if err != nil {
			break
		}

		steps++
	}

	return steps
}
