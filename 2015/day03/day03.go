package main

import (
	"fmt"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")
	fmt.Println("part1:", part1(lines[0]))
	fmt.Println("part2:", part2(lines[0]))
}

func part1(directions string) int {
	g := util.IntGrid2{}
	pos := util.Origin2()
	g.Set(pos, 1)

	for _, c := range directions {
		switch c {
		case '^':
			pos = pos.Offset(util.Point2{0, -1})
		case 'v':
			pos = pos.Offset(util.Point2{0, 1})
		case '<':
			pos = pos.Offset(util.Point2{-1, 0})
		case '>':
			pos = pos.Offset(util.Point2{1, 0})
		}
		g.Set(pos, g.Get(pos)+1)
	}

	return len(g)
}

func part2(directions string) int {
	g := util.IntGrid2{}
	pos := []util.Point2{util.Origin2(), util.Origin2()}
	cur := 0
	g.Set(util.Origin2(), 2)

	for _, c := range directions {
		switch c {
		case '^':
			pos[cur] = pos[cur].Offset(util.Point2{0, -1})
		case 'v':
			pos[cur] = pos[cur].Offset(util.Point2{0, 1})
		case '<':
			pos[cur] = pos[cur].Offset(util.Point2{-1, 0})
		case '>':
			pos[cur] = pos[cur].Offset(util.Point2{1, 0})
		}
		g.Set(pos[cur], g.Get(pos[cur])+1)
		cur = 1 - cur
	}

	return len(g)
}
