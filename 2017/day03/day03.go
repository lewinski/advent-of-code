package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.IntLines("input.txt")

	fmt.Println("part1:", part1(lines[0]))
	fmt.Println("part2:", part2(lines[0]))
}

func spiralStep(pos util.Point2) util.Point2 {
	up := util.Point2{0, 1}
	down := util.Point2{0, -1}
	left := util.Point2{-1, 0}
	right := util.Point2{1, 0}

	if pos[0] >= 0 && pos[0] == -pos[1] {
		// move out when in lower right quadrant and on y = -x line
		return pos.Offset(right)
	} else if pos[1] < pos[0] && pos[1] >= -pos[0] {
		// up when   y < x  && y >= -x
		return pos.Offset(up)
	} else if pos[1] >= pos[0] && pos[1] > -pos[0] {
		// left when y >= x && y > -x
		return pos.Offset(left)
	} else if pos[1] > pos[0] && pos[1] <= -pos[0] {
		// down when y > x  && y <= -x
		return pos.Offset(down)
	} else if pos[1] <= pos[0] && pos[1] < -pos[0] {
		// right when y <= x && y < -x
		return pos.Offset(right)
	}

	panic("can't happen")
}

func part1(input int) int {
	g := make(util.IntGrid2)

	cur := util.Origin2()
	num := 1

	for num != input {
		g.Set(cur, num)
		num++
		cur = spiralStep(cur)
	}

	return util.IAbs(cur[0]) + util.IAbs(cur[1])
}

func part2(input int) int {
	g := make(util.IntGrid2)

	cur := util.Origin2()
	g.Set(cur, 1)

	for {
		cur = spiralStep(cur)

		sum := 0
		for _, p := range cur.Around() {
			sum += g.Get(p)
		}
		g.Set(cur, sum)

		if sum > input {
			return sum
		}
	}
}
