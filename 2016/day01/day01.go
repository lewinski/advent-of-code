package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")[0]

	directions := strings.Split(input, ", ")

	pos := util.Origin2()
	facing := util.Point2{0, 1}
	visited := util.IntGrid2{}
	var part2 *util.Point2

	for _, d := range directions {
		if d[0] == 'L' {
			facing = left(facing)
		} else {
			facing = right(facing)
		}

		steps := util.MustAtoi(d[1:])
		for i := 0; i < steps; i++ {
			if part2 == nil && visited.Contains(pos) {
				part2 = &util.Point2{pos[0], pos[1]}
			}
			visited.Set(pos, 1)
			pos = pos.Offset(facing)
		}
	}

	fmt.Println("part1:", util.IAbs(pos[0])+util.IAbs(pos[1]))
	fmt.Println("part2:", util.IAbs(part2[0])+util.IAbs(part2[1]))
}

func right(p util.Point2) util.Point2 {
	return util.Point2{p[1], -p[0]}
}

func left(p util.Point2) util.Point2 {
	return util.Point2{-p[1], p[0]}
}
